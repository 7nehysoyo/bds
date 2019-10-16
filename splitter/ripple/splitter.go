package ripple

import (
	"fmt"
	"github.com/go-xorm/xorm"
	"github.com/jdcloud-bds/bds/common/httputils"
	"github.com/jdcloud-bds/bds/common/jsonrpc"
	"github.com/jdcloud-bds/bds/common/kafka"
	"github.com/jdcloud-bds/bds/common/log"
	"github.com/jdcloud-bds/bds/service"
	model "github.com/jdcloud-bds/bds/service/model/ripple"
	"github.com/xeipuuv/gojsonschema"
	"strconv"
	"strings"
	"time"
)

type SplitterConfig struct {
	Engine                     *xorm.Engine
	Consumer                   *kafka.ConsumerGroup
	Topic                      string
	DatabaseEnable             bool
	MaxBatchBlock              int
	Endpoint                   string
	User                       string
	Password                   string
	JSONSchemaFile             string
	JSONSchemaValidationEnable bool
	DatabaseWorkerNumber       int
	DatabaseWorkerBuffer       int
}

type XRPSplitter struct {
	cfg                           *SplitterConfig
	remoteHandler                 *rpcHandler
	cronWorker                    *CronWorker
	jsonSchemaLoader              gojsonschema.JSONLoader
	missedBlockList               map[int64]bool
	latestSaveDataTimestamp       time.Time
	latestReceiveMessageTimestamp time.Time
	databaseWorkerChan            chan *XRPLedgerData
	databaseWorkerStopChan        chan bool
}

func NewSplitter(cfg *SplitterConfig) (*XRPSplitter, error) {
	var err error
	s := new(XRPSplitter)
	s.cfg = cfg
	s.databaseWorkerChan = make(chan *XRPLedgerData, cfg.DatabaseWorkerBuffer)
	s.databaseWorkerStopChan = make(chan bool, 0)
	s.missedBlockList = make(map[int64]bool, 0)
	httpClient := httputils.NewRestClientWithBasicAuth(s.cfg.User, s.cfg.Password)
	s.remoteHandler, err = newRPCHandler(jsonrpc.New(httpClient, s.cfg.Endpoint))
	if err != nil {
		log.DetailError(err)
	}

	if s.cfg.JSONSchemaValidationEnable {
		f := fmt.Sprintf("file://%s", s.cfg.JSONSchemaFile)
		s.jsonSchemaLoader = gojsonschema.NewReferenceLoader(f)
	}

	s.cronWorker = NewCronWorker(s)
	//err = s.cronWorker.Prepare()
	//if err != nil {
	//	log.DetailError(err)
	//	return nil, err
	//}

	return s, nil
}

func (s *XRPSplitter) Start() {
	err := s.cronWorker.Start()
	if err != nil {
		log.Error("splitter ripple: cron worker start error")
		log.DetailError(err)
		return
	}

	err = s.cfg.Consumer.Start(s.cfg.Topic)
	if err != nil {
		log.Error("splitter ripple: consumer start error")
		log.DetailError(err)
		return
	}

	for i := 0; i < s.cfg.DatabaseWorkerNumber; i++ {
		go s.databaseWorker(i)
	}

	log.Debug("splitter ripple: consumer start topic %s", s.cfg.Topic)
	log.Debug("splitter ripple: database enable is %v", s.cfg.DatabaseEnable)

	for {
		select {
		case message := <-s.cfg.Consumer.MessageChannel():
			stats.Add(MetricReceiveMessages, 1)
			s.latestReceiveMessageTimestamp = time.Now()

		START:
			if s.cfg.JSONSchemaValidationEnable {
				ok, err := s.jsonSchemaValid(string(message.Data))
				if err != nil {
					log.Error("splitter ripple: json schema valid error")
				}
				if !ok {
					log.Warn("splitter ripple: json schema valid failed")
				}
			}

			data, err := ParseLedger(string(message.Data))
			if err != nil {
				stats.Add(MetricParseDataError, 1)
				log.Error("splitter ripple: ledger parse error, retry after 5s")
				log.DetailError(err)
				time.Sleep(time.Second * 5)
				goto START
			}

			exist, err := s.IsExistingLedger(data)
			if err != nil {
				stats.Add(MetricParseDataError, 1)
				log.Error("splitter ripple: check ledger exist error, retry after 2s")
				log.DetailError(err)
				time.Sleep(time.Second * 2)
				goto START
			}

			if s.cfg.DatabaseEnable && !exist {
				s.databaseWorkerChan <- data
				s.cfg.Consumer.MarkOffset(message)
			}
		}
	}
}

func (s *XRPSplitter) Stop() {
	s.cronWorker.Stop()
}

func (s *XRPSplitter) IsExistingLedger(cur *XRPLedgerData) (bool, error) {
	db := service.NewDatabase(s.cfg.Engine)
	ledgers := make([]*model.Ledger, 0)
	err := db.Where("ledger_index = ?", cur.Ledger.LedgerIndex).Find(&ledgers)
	if err != nil {
		log.DetailError(err)
		return false, err
	}

	if len(ledgers) == 0 {
		//log.Warn("splitter ripple: can not find current ledger %d", cur.Ledger.LedgerIndex)
		return false, nil
	}

	return true, nil
}

func (s *XRPSplitter) LedgerInsert(data *XRPLedgerData) error {
	startTime := time.Now()
	tx := service.NewTransaction(s.cfg.Engine)
	defer tx.Close()

	err := tx.Begin()
	if err != nil {
		_ = tx.Rollback()
		log.DetailError(err)
		stats.Add(MetricDatabaseRollback, 1)
		return err
	}

	var affected int64
	ledgers := make([]*model.Ledger, 0)
	ledgers = append(ledgers, data.Ledger)
	affected, err = tx.BatchInsert(ledgers)
	if err != nil {
		_ = tx.Rollback()
		log.DetailError(err)
		stats.Add(MetricDatabaseRollback, 1)
		return err
	}
	log.Debug("splitter ripple: ledger write %d rows", affected)

	if len(data.AccountSets) != 0 {
		affected, err = tx.BatchInsert(data.AccountSets)
		if err != nil {
			_ = tx.Rollback()
			log.DetailError(err)
			stats.Add(MetricDatabaseRollback, 1)
			return err
		}
		log.Debug("splitter ripple: transaction AccountSets write %d rows", affected)
	}

	if len(data.DepositPreauths) != 0 {
		affected, err = tx.BatchInsert(data.DepositPreauths)
		if err != nil {
			_ = tx.Rollback()
			log.DetailError(err)
			stats.Add(MetricDatabaseRollback, 1)
			return err
		}
		log.Debug("splitter ripple: transaction DepositPreauths write %d rows", affected)
	}

	if len(data.EscrowCancels) != 0 {
		affected, err = tx.BatchInsert(data.EscrowCancels)
		if err != nil {
			_ = tx.Rollback()
			log.DetailError(err)
			stats.Add(MetricDatabaseRollback, 1)
			return err
		}
		log.Debug("splitter ripple: transaction EscrowCancels write %d rows", affected)
	}

	if len(data.EscrowCreates) != 0 {
		affected, err = tx.BatchInsert(data.EscrowCreates)
		if err != nil {
			_ = tx.Rollback()
			log.DetailError(err)
			stats.Add(MetricDatabaseRollback, 1)
			return err
		}
		log.Debug("splitter ripple: transaction EscrowCreates write %d rows", affected)
	}

	if len(data.EscrowFinishes) != 0 {
		affected, err = tx.BatchInsert(data.EscrowFinishes)
		if err != nil {
			_ = tx.Rollback()
			log.DetailError(err)
			stats.Add(MetricDatabaseRollback, 1)
			return err
		}
		log.Debug("splitter ripple: transaction EscrowFinishes write %d rows", affected)
	}

	if len(data.OfferCancels) != 0 {
		affected, err = tx.BatchInsert(data.OfferCancels)
		if err != nil {
			_ = tx.Rollback()
			log.DetailError(err)
			stats.Add(MetricDatabaseRollback, 1)
			return err
		}
		log.Debug("splitter ripple: transaction OfferCancels write %d rows", affected)
	}

	if len(data.OfferCreates) != 0 {
		affected, err = tx.BatchInsert(data.OfferCreates)
		if err != nil {
			_ = tx.Rollback()
			log.DetailError(err)
			stats.Add(MetricDatabaseRollback, 1)
			return err
		}
		log.Debug("splitter ripple: transaction OfferCreates write %d rows", affected)
	}

	if len(data.Payments) != 0 {
		affected, err = tx.BatchInsert(data.Payments)
		if err != nil {
			_ = tx.Rollback()
			log.DetailError(err)
			stats.Add(MetricDatabaseRollback, 1)
			return err
		}
		log.Debug("splitter ripple: transaction Payments write %d rows", affected)
	}

	if len(data.PaymentChannelClaims) != 0 {
		affected, err = tx.BatchInsert(data.PaymentChannelClaims)
		if err != nil {
			_ = tx.Rollback()
			log.DetailError(err)
			stats.Add(MetricDatabaseRollback, 1)
			return err
		}
		log.Debug("splitter ripple: transaction PaymentChannelClaims write %d rows", affected)
	}

	if len(data.PaymentChannelCreates) != 0 {
		affected, err = tx.BatchInsert(data.PaymentChannelCreates)
		if err != nil {
			_ = tx.Rollback()
			log.DetailError(err)
			stats.Add(MetricDatabaseRollback, 1)
			return err
		}
		log.Debug("splitter ripple: transaction PaymentChannelCreates write %d rows", affected)
	}

	if len(data.PaymentChannelFunds) != 0 {
		affected, err = tx.BatchInsert(data.PaymentChannelFunds)
		if err != nil {
			_ = tx.Rollback()
			log.DetailError(err)
			stats.Add(MetricDatabaseRollback, 1)
			return err
		}
		log.Debug("splitter ripple: transaction PaymentChannelFunds write %d rows", affected)
	}

	if len(data.SetRegularKeys) != 0 {
		affected, err = tx.BatchInsert(data.SetRegularKeys)
		if err != nil {
			_ = tx.Rollback()
			log.DetailError(err)
			stats.Add(MetricDatabaseRollback, 1)
			return err
		}
		log.Debug("splitter ripple: transaction SetRegularKeys write %d rows", affected)
	}

	if len(data.SignerListSets) != 0 {
		affected, err = tx.BatchInsert(data.SignerListSets)
		if err != nil {
			_ = tx.Rollback()
			log.DetailError(err)
			stats.Add(MetricDatabaseRollback, 1)
			return err
		}
		log.Debug("splitter ripple: transaction SignerListSets write %d rows", affected)
	}

	if len(data.TrustSets) != 0 {
		affected, err = tx.BatchInsert(data.TrustSets)
		if err != nil {
			_ = tx.Rollback()
			log.DetailError(err)
			stats.Add(MetricDatabaseRollback, 1)
			return err
		}
		log.Debug("splitter ripple: transaction TrustSets write %d rows", affected)
	}

	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		log.DetailError(err)
		return err
	}
	tx.Close()
	stats.Add(MetricDatabaseCommit, 1)
	elapsedTime := time.Now().Sub(startTime)
	s.latestSaveDataTimestamp = time.Now()
	log.Debug("splitter ripple: ledger data %d write done elapsed: %s", data.Ledger.LedgerIndex, elapsedTime.String())
	return nil
}

func (s *XRPSplitter) revertLedger(ledgerIndex int64) error {
	tx := service.NewTransaction(s.cfg.Engine)
	defer tx.Close()

	err := tx.Begin()
	if err != nil {
		_ = tx.Rollback()
		log.DetailError(err)
		stats.Add(MetricDatabaseRollback, 1)
		return err
	}
	sql := fmt.Sprintf("delete from ripple_ledger where ledger_index = %d", ledgerIndex)
	_, err = tx.Exec(sql)
	if err != nil {
		_ = tx.Rollback()
		log.DetailError(err)
		stats.Add(MetricDatabaseRollback, 1)
		return err
	}
	sql = fmt.Sprintf("delete from ripple_ac where ledger_index = %d", ledgerIndex)
	_, err = tx.Exec(sql)
	if err != nil {
		_ = tx.Rollback()
		log.DetailError(err)
		stats.Add(MetricDatabaseRollback, 1)
		return err
	}
	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		log.DetailError(err)
		return err
	}
	tx.Close()
	return nil
}

func (s *XRPSplitter) CheckMissedLedger() ([]int64, error) {
	missedList := make([]int64, 0)

	db := service.NewDatabase(s.cfg.Engine)
	sql := fmt.Sprintf("SELECT ledger_index FROM ripple_ledger ORDER BY ledger_index ASC")
	data, err := db.QueryString(sql)
	if err != nil {
		return nil, err
	}

	ledgerList := make([]*model.Ledger, 0)
	for _, v := range data {
		ledger := new(model.Ledger)
		tmp := v["ledger_index"]
		ledgerIndex, err := strconv.ParseInt(tmp, 10, 64)
		if err != nil {
			return nil, err
		}
		ledger.LedgerIndex = ledgerIndex
		ledgerList = append(ledgerList, ledger)
	}

	if len(ledgerList) > 0 {
		checkList := make(map[int64]bool, 0)
		for _, b := range ledgerList {
			checkList[b.LedgerIndex] = true
		}

		for i := int64(32570); i <= ledgerList[len(ledgerList)-1].LedgerIndex; i++ {
			if _, ok := checkList[i]; !ok {
				missedList = append(missedList, i)
			}
		}
	}

	return missedList, nil
}

func (s *XRPSplitter) jsonSchemaValid(data string) (bool, error) {
	startTime := time.Now()
	dataLoader := gojsonschema.NewStringLoader(data)
	result, err := gojsonschema.Validate(s.jsonSchemaLoader, dataLoader)
	if err != nil {
		log.Error("splitter ripple: json schema validation error")
		log.DetailError(err)
		return false, err
	}
	if !result.Valid() {
		for _, err := range result.Errors() {
			log.Error("splitter ripple: data invalid %s", strings.ToLower(err.String()))
			return false, nil
		}
		stats.Add(MetricVaildationError, 1)
	} else {
		stats.Add(MetricVaildationSuccess, 1)
	}
	elaspedTime := time.Now().Sub(startTime)
	log.Debug("splitter ripple: json schema validation elapsed %s", elaspedTime)
	return true, nil
}

func (s *XRPSplitter) databaseWorker(i int) {
	log.Info("splitter ripple: starting database worker %d", i)
	for {
		select {
		case data := <-s.databaseWorkerChan:
			err := s.LedgerInsert(data)
			if err != nil {
				log.Error("splitter ripple: ledger data %d insert error, retry after 5s", data.Ledger.LedgerIndex)
				log.DetailError(err)
			}
		case stop := <-s.databaseWorkerStopChan:
			if stop {
				msg := fmt.Sprintf("splitter ripple: database worker %d stopped", i)
				log.Info("splitter ripple: ", msg)
				return
			}
		}
	}
}
