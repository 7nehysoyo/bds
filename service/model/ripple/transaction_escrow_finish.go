package ripple

type EscrowFinish struct {
	TransactionCommonFields `xorm:"extends"`

	Owner         string `xorm:"owner varchar(68) null"`
	OfferSequence int64  `xorm:"offer_sequence bigint null"`
	Condition     string `xorm:"condition varchar(512) null"`
	Fulfillment   string `xorm:"fulfillment varchar(512) null"`
}

func (t EscrowFinish) TableName() string {
	return tableName("transaction_escrow_finish")
}
