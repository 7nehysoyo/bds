package ripple

import (
	"github.com/jdcloud-bds/bds/common/math"
)

type Ledger struct {
	ID                  int64                `xorm:"id bigint autoincr pk"`
	Accepted            int                  `xorm:"accepted tinyint notnull"`
	AccountHash         string               `xorm:"account_hash varchar(128) notnull"`
	CloseFlags          int                  `xorm:"close_flags tinyint notnull"`
	Timestamp           int64                `xorm:"timestamp int notnull index"`
	CloseTime           int64                `xorm:"close_time int notnull"`
	CloseTimeHuman      string               `xorm:"close_time_human varchar(56) notnull"`
	CloseTimeResolution int                  `xorm:"close_time_resolution int notnull"`
	Closed              int                  `xorm:"closed tinyint notnull"`
	Hash                string               `xorm:"hash varchar(128) not null"`
	LedgerHash          string               `xorm:"ledger_hash varchar(128) notnull"`
	LedgerIndex         int64                `xorm:"ledger_index bigint unique notnull index"`
	ParentCloseTime     int64                `xorm:"parent_close_time int notnull"`
	ParentHash          string               `xorm:"parent_hash varchar(128) notnull"`
	SeqNum              int64                `xorm:"seq_num bigint notnull"`
	TotalCoins          math.HexOrDecimal256 `xorm:"total_coins decimal(38,0) null"`
	TransactionHash     string               `xorm:"transaction_hash varchar(128) notnull"`
	TransactionLength   int                  `xorm:"tx_len int notnull"`
}

func (t Ledger) TableName() string {
	return tableName("ledger")
}
