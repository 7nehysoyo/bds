package ripple

import (
	"github.com/jdcloud-bds/bds/common/cron"
	"github.com/jdcloud-bds/bds/common/log"
	"github.com/jdcloud-bds/bds/config"
)

type CronWorker struct {
	splitter *XRPSplitter
	crontab  *cron.Cron
}

func NewCronWorker(splitter *XRPSplitter) *CronWorker {
	worker := new(CronWorker)
	worker.splitter = splitter
	worker.crontab = cron.New()
	return worker
}

func (w *CronWorker) Prepare() error {
	jobList := []WorkerJob{
		newUpdateMetaDataJob(w.splitter),
	}

	for _, job := range jobList {
		log.Debug("worker ripple: prepare %s", job.Name())
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

	expr = config.SplitterConfig.CronXRPSetting.UpdateMetaExpr
	if len(expr) != 0 {
		job = newUpdateMetaDataJob(w.splitter)
		err = w.crontab.AddJob(job.Name(), expr, job)
		if err != nil {
			return err
		}
		stats.Add(MetricCronWorkerJob, 1)
		log.Info("worker ripple: add job %s", job.Name())
	}

	expr = config.SplitterConfig.CronXRPSetting.GetBatchLedgerExpr
	if len(expr) != 0 {
		job = newGetBatchLedgerJob(w.splitter)
		err = w.crontab.AddJob(job.Name(), expr, job)
		if err != nil {
			return err
		}
		stats.Add(MetricCronWorkerJob, 1)
		log.Info("worker ripple: add job %s", job.Name())
	}

	w.crontab.Start()
	return nil
}

func (w *CronWorker) Stop() {
	w.crontab.Stop()
}
