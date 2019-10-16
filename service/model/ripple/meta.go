package ripple

import (
	"fmt"
	"time"
)

const (
	TablePrefix = "ripple"
)

type Meta struct {
	ID          int64     `xorm:"id bigint autoincr pk"`
	Name        string    `xorm:"name varchar(255) notnull unique"`
	LastID      int64     `xorm:"last_id bigint notnull"`
	Count       int64     `xorm:"count bigint notnull"`
	CreatedTime time.Time `xorm:"created_time created notnull"`
	UpdatedTime time.Time `xorm:"updated_time updated notnull"`
}

func (t Meta) TableName() string {
	return tableName("meta")
}

func tableName(s string) string {
	if len(TablePrefix) == 0 {
		return s
	}
	return fmt.Sprintf("%s_%s", TablePrefix, s)
}

type TransactionCommonFields struct {
	ID int64 `xorm:"id bigint autoincr pk"`
	// The type of transaction. Valid types include: Payment, OfferCreate, OfferCancel, TrustSet,
	// AccountSet, SetRegularKey, SignerListSet, EscrowCreate, EscrowFinish, EscrowCancel,
	// PaymentChannelCreate, PaymentChannelFund, PaymentChannelClaim, and DepositPreauth.
	TransactionType    string `xorm:"transaction_type char(32) notnull index"`
	Account            string `xorm:"account varchar(68) notnull index"`
	Hash               string `xorm:"hash varchar(128) notnull index"`
	LedgerIndex        int64  `xorm:"ledger_index bigint index"`
	Timestamp          int64  `xorm:"timestamp int notnull index"`
	Fee                int64  `xorm:"fee bigint notnull"` //in drops
	Sequence           int64  `xorm:"sequence bigint notnull"`
	AccountTxnID       string `xorm:"account_txn_id varchar(128) null"`
	Flags              int64  `xorm:"flags bigint null"`
	LastLedgerSequence int64  `xorm:"last_ledger_sequence bigint null"`
	Memos              string `xorm:"memos text null"`
	Signers            string `xorm:"signers text null"`
	SourceTag          int64  `xorm:"source_tag bigint null"`
	SigningPubKey      string `xorm:"signing_pub_key varchar(128) null"`
	TxnSignature       string `xorm:"txn_signature varchar(256) null"`

	//metadata
	TransactionResult string `xorm:"transaction_result varchar(32) null"`
	TransactionIndex  int    `xorm:"transaction_index bigint null"`
	AffectedNodes     string `xorm:"affected_nodes text null"`
	DeliveredAmount   string `xorm:"delivered_amount text null"`
}
