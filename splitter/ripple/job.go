package ripple

import (
	"fmt"
	"github.com/jdcloud-bds/bds/common/log"
	"github.com/jdcloud-bds/bds/service"
	model "github.com/jdcloud-bds/bds/service/model/xrp"
	"strconv"
	"time"
)

type WorkerJob interface {
	Run()
	Name() string
	run() error
}

type updateMetaDataJob struct {
	splitter *XRPSplitter
	name     string
}

func newUpdateMetaDataJob(splitter *XRPSplitter) *updateMetaDataJob {
	j := new(updateMetaDataJob)
	j.splitter = splitter
	j.name = "update meta data"
	return j
}

func (j *updateMetaDataJob) Run() {
	_ = j.run()
}

func (j *updateMetaDataJob) Name() string {
	return j.name
}

func (j *updateMetaDataJob) run() error {
	startTime := time.Now()
	db := service.NewDatabase(j.splitter.cfg.Engine)
	metas := make([]*model.Meta, 0)
	err := db.Find(&metas)
	if err != nil {
		log.Error("ripple job : '%s' get table list from meta error", j.name)
		return err
	}

	for _, meta := range metas {
		cond := new(model.Meta)
		cond.Name = meta.Name
		data := new(model.Meta)

		var countSql string
		if j.splitter.cfg.Engine.DriverName() == "mssql" {
			countSql = fmt.Sprintf("SELECT b.rows AS count FROM sysobjects a INNER JOIN sysindexes b ON a.id = b.id WHERE a.type = 'u' AND b.indid in (0,1) AND a.name='%s'", meta.Name)
		} else {
			countSql = fmt.Sprintf("SELECT COUNT(1) FROM `%s`", meta.Name)
		}
		result, err := db.QueryString(countSql)
		if err != nil {
			log.Error("ripple job : %s get table %s count from meta error", j.name, meta.Name)
			log.DetailError(err)
			continue
		}
		if len(result) == 0 {
			continue
		}
		count, _ := strconv.ParseInt(result[0]["count"], 10, 64)

		sql := db.Table(meta.Name).Cols("id").Desc("id").Limit(1, 0)
		result, err = sql.QueryString()
		if err != nil {
			log.Error("ripple job : '%s' get table %s id from meta error", j.name, meta.Name)
			log.DetailError(err)
			continue
		}
		for _, v := range result {
			id, _ := strconv.ParseInt(v["id"], 10, 64)
			data.LastID = id
			data.Count = count
			_, err = db.Update(data, cond)
			if err != nil {
				log.Error("ripple job : '%s' update table %s meta error", j.name, meta.Name)
				log.DetailError(err)
				continue
			}

		}
	}
	stats.Add(MetricCronWorkerJobUpdateMetaData, 1)
	elapsedTime := time.Now().Sub(startTime)
	log.Debug("ripple job : '%s' elapsed time %s", j.name, elapsedTime.String())
	return nil
}

type getBatchLedgerJob struct {
	splitter *XRPSplitter
	name     string
}

func newGetBatchLedgerJob(splitter *XRPSplitter) *getBatchLedgerJob {
	j := new(getBatchLedgerJob)
	j.splitter = splitter
	j.name = "'get batch ledger'"
	return j
}

func (j *getBatchLedgerJob) Run() {
	_ = j.run()
}

func (j *getBatchLedgerJob) Name() string {
	return j.name
}

func (j *getBatchLedgerJob) run() error {
	startTime := time.Now()
	db := service.NewDatabase(j.splitter.cfg.Engine)
	totalCompleteLedgers, err := j.splitter.remoteHandler.GetCompleteLedgers()
	if err != nil {
		log.Error("ripple job : %s get closed ledgers error", j.name)
		log.DetailError("ripple job error : %s", err.Error())
		return err
	}
	batchSize := 1000

	ledgerList := make(map[int64]bool, 0)
	sql := "select ledger_index from ripple_ledger order by ledger_index asc"
	result, err := db.QueryString(sql)
	if err != nil {
		log.Error("ripple job : get ledger_index from database error")
		log.DetailError(err)
		return err
	}
	for _, v := range result {
		ledgerIndex, _ := strconv.ParseInt(v["ledger_index"], 10, 64)
		ledgerList[ledgerIndex] = true
	}
	for _, cl := range totalCompleteLedgers {
		missedLedger := make(map[int64]bool, 0)
		i := 0
		start := int64(0)
		end := int64(0)
		flag := false
		for k := cl.startLedger; k <= cl.endLedger; k++ {
			if _, ok := ledgerList[k]; !ok {
				if !flag {
					start = k
					end = k
					flag = true
				} else {
					end = k
				}
				missedLedger[k] = true
				i++
			} else if flag {
				break
			}

			if i >= batchSize {
				break
			}
		}
		if len(missedLedger) > 0 && end >= start {
			log.Info("splitter ripple: send batch ledger from %d to %d.", start, end)
			err = j.splitter.remoteHandler.SendBatchLedger(start, end)
			if err != nil {
				log.Error("splitter ripple: send batch ledger error: %s", err.Error())
				log.DetailError(err)
				return err
			}
		}
	}

	elapsedTime := time.Now().Sub(startTime)
	log.Debug("ripple job : '%s' elapsed time %s", j.name, elapsedTime.String())
	return nil
}
