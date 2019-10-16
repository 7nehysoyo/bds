package ripple

type AccountSet struct {
	TransactionCommonFields `xorm:"extends"`

	ClearFlag    int64  `xorm:"clear_flags bigint null"`
	Domain       string `xorm:"domain varchar(1024) null"`
	EmailHash    string `xorm:"email_hash char(128) null"`
	MessageKey   string `xorm:"message_key varchar(128) null"`
	SetFlag      int64  `xorm:"set_flag bigint null"`
	TransferRate int64  `xorm:"transfer_rate bigint null"`
	TickSize     int    `xorm:"tick_size int null"`
}

func (t AccountSet) TableName() string {
	return tableName("transaction_account_set")
}
