package ripple

type EscrowCreate struct {
	TransactionCommonFields `xorm:"extends"`

	Amount         string `xorm:"amount text notnull"`
	Destination    string `xorm:"destination varchar(68) notnull"`
	CancelAfter    int64  `xorm:"cancel_after bigint null"`
	FinishAfter    int64  `xorm:"finish_after bigint null"`
	Condition      string `xorm:"condition varchar(512) null"`
	DestinationTag int64  `xorm:"destination_tag bigint null"`
}

func (t EscrowCreate) TableName() string {
	return tableName("transaction_escrow_create")
}
