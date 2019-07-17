package eth

import (
	"github.com/jdcloud-bds/bds/common/cuckoofilter"
	"github.com/jdcloud-bds/bds/common/metric"
	model "github.com/jdcloud-bds/bds/service/model/eth"
	"math/big"
	"sync"
)

const (
	MetricReceiveMessages                       = "receive_messages"
	MetricParseDataError                        = "parse_data_error"
	MetricVaildationSuccess                     = "validation_success"
	MetricVaildationError                       = "validation_error"
	MetricDatabaseRollback                      = "database_rollback"
	MetricDatabaseCommit                        = "database_commit"
	MetricCronWorkerJob                         = "cron_worker_job"
	MetricCronWorkerJobUpdateMetaData           = "cron_worker_job_update_meta_data"
	MetricCronWorkerJobGetBatchBlock            = "cron_worker_job_get_batch_block"
	MetricCronWorkerJobRefreshContractAddresses = "cron_worker_job_refresh_contract_addresses"
	MetricCronWorkerJobRefreshPoolName          = "cron_worker_job_refresh_pool_name"
	MetricRPCCall                               = "rpc_call"
	MetricRevertBlock                           = "revert_block"

	AccountTypeNormal   = 0
	AccountTypeContract = 1
	AccountTypeMiner    = 2

	TransactionTypeNormal   = 0
	TransactionTypeContract = 1
)

var (
	stats                 = metric.NewMap("eth")
	contractAddressFilter = cuckoofilter.New()
	poolNameMap           = new(sync.Map)
	maxBigNumber, _       = new(big.Int).SetString("100000000000000000000000000000000000000", 10)
	defaultBigNumber, _   = new(big.Int).SetString("-1", 10)
)

type ETHBlockData struct {
	Block                *model.Block
	Uncles               []*model.Uncle
	Transactions         []*model.Transaction
	InternalTransactions []*model.InternalTransaction
	TokenTransactions    []*model.TokenTransaction
	Tokens               []*model.Token
	ENSes                []*model.ENS
	Accounts             []*model.Account
	TokenAccounts        []*model.TokenAccount
}
