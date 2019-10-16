package ripple

import (
	"github.com/jdcloud-bds/bds/common/metric"
	model "github.com/jdcloud-bds/bds/service/model/ripple"
	"math/big"
)

const (
	MetricReceiveMessages             = "receive_messages"
	MetricParseDataError              = "parse_data_error"
	MetricVaildationSuccess           = "validation_success"
	MetricVaildationError             = "validation_error"
	MetricDatabaseRollback            = "database_rollback"
	MetricDatabaseCommit              = "database_commit"
	MetricCronWorkerJob               = "cron_worker_job"
	MetricCronWorkerJobUpdateMetaData = "cron_worker_job_update_meta_data"
	MetricCronWorkerJobGetBatchLedger = "cron_worker_job_get_batch_ledger"
	MetricRPCCall                     = "rpc_call"

	ParamStartLedgerIndex = "start_ledger_index"
	ParamEndLedgerIndex   = "end_ledger_index"
)

const (
	AccountSet           = "AccountSet"
	DepositPreauth       = "DepositPreauth"
	EscrowCancel         = "EscrowCancel"
	EscrowCreate         = "EscrowCreate"
	EscrowFinish         = "EscrowFinish"
	OfferCancel          = "OfferCancel"
	OfferCreate          = "OfferCreate"
	Payment              = "Payment"
	PaymentChannelClaim  = "PaymentChannelClaim"
	PaymentChannelCreate = "PaymentChannelCreate"
	PaymentChannelFund   = "PaymentChannelFund"
	SetRegularKey        = "SetRegularKey"
	SignerListSet        = "SignerListSet"
	TrustSet             = "TrustSet"
)

var (
	stats               = metric.NewMap("ripple")
	maxBigNumber, _     = new(big.Int).SetString("100000000000000000000000000000000000000", 10)
	defaultBigNumber, _ = new(big.Int).SetString("-1", 10)
)

type XRPLedgerData struct {
	Ledger                *model.Ledger
	AccountSets           []*model.AccountSet
	DepositPreauths       []*model.DepositPreauth
	EscrowCancels         []*model.EscrowCancel
	EscrowCreates         []*model.EscrowCreate
	EscrowFinishes        []*model.EscrowFinish
	OfferCancels          []*model.OfferCancel
	OfferCreates          []*model.OfferCreate
	Payments              []*model.Payment
	PaymentChannelClaims  []*model.PaymentChannelClaim
	PaymentChannelCreates []*model.PaymentChannelCreate
	PaymentChannelFunds   []*model.PaymentChannelFund
	SetRegularKeys        []*model.SetRegularKey
	SignerListSets        []*model.SignerListSet
	TrustSets             []*model.TrustSet
}
