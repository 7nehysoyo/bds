package eos

import (
	"github.com/jdcloud-bds/bds/common/cron"
	"github.com/jdcloud-bds/bds/common/log"
	"github.com/jdcloud-bds/bds/config"
)

type CronWorker struct {
	splitter *EOSSplitter
	crontab  *cron.Cron
}

func NewCronWorker(splitter *EOSSplitter) *CronWorker {
	worker := new(CronWorker)
	worker.splitter = splitter
	worker.crontab = cron.New()
	return worker
}

func (w *CronWorker) Prepare() error {
	jobList := []WorkerJob{
		newUpdateMetaDataJob(w.splitter),
		//newRefreshContractAddressesJob(w.splitter),
		//newRefreshPoolNameJob(w.splitter),
	}

	for _, job := range jobList {
		log.Debug("worker eos: prepare %s", job.Name())
		err := job.run()
		if err != nil {
			return err
		}
	}

	return nil
}

func (w *CronWorker) Start() error {
	var err error
	var expr string
	var job WorkerJob

	expr = config.SplitterConfig.CronEOSSetting.UpdateMetaExpr
	if len(expr) != 0 {
		job = newUpdateMetaDataJob(w.splitter)
		err = w.crontab.AddJob(job.Name(), expr, job)
		if err != nil {
			return err
		}
		stats.Add(MetricCronWorkerJob, 1)
		log.Debug("worker eos: add job %s", job.Name())
	}

	w.crontab.Start()
	return nil
}

func (w *CronWorker) Stop() {
	w.crontab.Stop()
}
