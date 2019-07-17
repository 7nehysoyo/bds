package eth

import (
	"fmt"
	"github.com/jdcloud-bds/bds/common/cuckoofilter"
	"github.com/jdcloud-bds/bds/common/log"
	"github.com/jdcloud-bds/bds/service"
	model "github.com/jdcloud-bds/bds/service/model/eth"
	"strconv"
	"strings"
	"time"
)

type WorkerJob interface {
	Run()
	Name() string
	run() error
}

type updateMetaDataJob struct {
	splitter *ETHSplitter
	name     string
}

func newUpdateMetaDataJob(splitter *ETHSplitter) *updateMetaDataJob {
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
		log.Error("worker eth: job '%s' get table list from meta error", j.name)
		return err
	}
	//update meta table
	for _, meta := range metas {
		cond := new(model.Meta)
		cond.Name = meta.Name
		data := new(model.Meta)
		//count table size
		var countSql string
		if j.splitter.cfg.Engine.DriverName() == "mssql" {
			countSql = fmt.Sprintf("SELECT b.rows AS count FROM sysobjects a INNER JOIN sysindexes b ON a.id = b.id WHERE a.type = 'u' AND b.indid in (0,1) AND a.name='%s'", meta.Name)
		} else {
			countSql = fmt.Sprintf("SELECT COUNT(1) FROM `%s`", meta.Name)
		}
		result, err := db.QueryString(countSql)
		if err != nil {
			log.Error("worker eth: job %s get table %s count from meta error", j.name, meta.Name)
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
			log.Error("worker eth: job '%s' get table %s id from meta error", j.name, meta.Name)
			log.DetailError(err)
			continue
		}
		//update meta
		for _, v := range result {
			id, _ := strconv.ParseInt(v["id"], 10, 64)
			data.LastID = id
			data.Count = count
			_, err = db.Update(data, cond)
			if err != nil {
				log.Error("worker eth: job '%s' update table %s meta error", j.name, meta.Name)
				log.DetailError(err)
				continue
			}

		}
	}
	stats.Add(MetricCronWorkerJobUpdateMetaData, 1)
	elaspedTime := time.Now().Sub(startTime)
	log.Debug("worker eth: job '%s' elasped time %s", j.name, elaspedTime.String())
	return nil
}

type getBatchBlockJob struct {
	splitter *ETHSplitter
	name     string
}

func newGetBatchBlockJob(splitter *ETHSplitter) *getBatchBlockJob {
	j := new(getBatchBlockJob)
	j.splitter = splitter
	j.name = "'get batch block'"
	return j
}

func (j *getBatchBlockJob) Run() {
	_ = j.run()
}

func (j *getBatchBlockJob) Name() string {
	return j.name
}

func (j *getBatchBlockJob) run() error {
	startTime := time.Now()
	db := service.NewDatabase(j.splitter.cfg.Engine)
	blocks := make([]*model.Block, 0)
	err := db.Desc("height").Limit(1).Find(&blocks)
	if err != nil {
		log.Error("worker eth: job '%s' get latest block error", j.name)
	}
	var start, end int64
	var old, free bool

	now := time.Now()
	if len(blocks) > 0 {
		start = blocks[0].Height + 1
		end = blocks[0].Height + int64(j.splitter.cfg.MaxBatchBlock)
		if (now.Unix() - blocks[0].Timestamp) > 40 {
			old = true
		} else {
			old = false
		}
	} else {
		start = 0
		end = int64(j.splitter.cfg.MaxBatchBlock)
		old = true
	}

	if now.Sub(j.splitter.latestSaveDataTimestamp).Seconds() > 15 && now.Sub(j.splitter.latestReceiveMessageTimestamp) > 5 {
		log.Warn("worker eth: job '%s' splitter eth no message received or no data processing", j.name)
		free = true
	}

	if free && old {
		log.Info("worker eth: job '%s' get block range from %d to %d", j.name, start, end)
		err = j.splitter.remoteHandler.SendBatchBlock(start, end)
		if err != nil {
			log.Error("worker eth: job '%s' rpc call error", j.name)
			log.DetailError(err)
		}

		stats.Add(MetricCronWorkerJobGetBatchBlock, 1)
	}

	elaspedTime := time.Now().Sub(startTime)
	log.Debug("worker eth: job '%s' elasped time %s", j.name, elaspedTime.String())
	return nil
}

type refreshContractAddressesJob struct {
	splitter *ETHSplitter
	name     string
}

func newRefreshContractAddressesJob(splitter *ETHSplitter) *refreshContractAddressesJob {
	j := new(refreshContractAddressesJob)
	j.splitter = splitter
	j.name = "refresh contract addresses"
	return j
}

func (j *refreshContractAddressesJob) Run() {
	_ = j.run()
}

func (j *refreshContractAddressesJob) Name() string {
	return j.name
}

func (j *refreshContractAddressesJob) run() error {
	startTime := time.Now()
	db := service.NewDatabase(j.splitter.cfg.Engine)
	//accounts := make([]*model.Account, 0)
	//err := db.Where("type = ? OR type = ?", 1, 11).Find(&accounts)
	//err := db.Where(builder.Eq{"type": 1}.Or(builder.Eq{"type": 11})).Find(&accounts)
	sql := fmt.Sprintf("SELECT address FROM eth_account WHERE type = 1 OR type = 11")
	result, err := db.QueryString(sql)
	if err != nil {
		log.DetailError(err)
		return err
	}

	if err != nil {
		log.Error("worker eth: job '%s' get contract account list from account error", j.name)
		return err
	}

	filter := cuckoofilter.NewCuckooFilter()
	for _, v := range result {
		address := strings.ToLower(v["address"])
		ok := filter.Insert([]byte(address))
		if !ok {
			log.Warn("worker eth: job '%s' insert address %s to filter error", j.name, address)
		}
	}
	oldCount := contractAddressFilter.Count()
	contractAddressFilter.Update(filter)
	log.Debug("worker eth: job '%s' contract addresses count %d:%d", j.name, oldCount, filter.Count())

	stats.Add(MetricCronWorkerJobRefreshContractAddresses, 1)
	elaspedTime := time.Now().Sub(startTime)
	log.Debug("worker eth: job '%s' account update elasped time %s", j.name, elaspedTime.String())
	return nil
}

type refreshPoolNameJob struct {
	splitter *ETHSplitter
	name     string
}

func newRefreshPoolNameJob(splitter *ETHSplitter) *refreshPoolNameJob {
	j := new(refreshPoolNameJob)
	j.splitter = splitter
	j.name = "refresh pool name"
	return j
}

func (j *refreshPoolNameJob) Run() {
	_ = j.run()
}

func (j *refreshPoolNameJob) Name() string {
	return j.name
}

func (j *refreshPoolNameJob) run() error {
	startTime := time.Now()
	db := service.NewDatabase(j.splitter.cfg.Engine)
	pools := make([]*model.MinerPoolAddress, 0)
	err := db.Find(&pools)
	if err != nil {
		log.Error("worker eth: job '%s' get pool name list from pool_name error", j.name)
		return err
	}

	for _, pool := range pools {
		poolNameMap.Store(pool.Address, pool.Name)
	}

	log.Debug("worker eth: job '%s' pool name map count %d", j.name, len(pools))

	stats.Add(MetricCronWorkerJobRefreshPoolName, 1)
	elaspedTime := time.Now().Sub(startTime)
	log.Debug("worker eth: job '%s' account update elasped time %s", j.name, elaspedTime.String())
	return nil
}
