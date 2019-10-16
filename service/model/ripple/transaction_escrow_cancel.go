package ripple

type EscrowCancel struct {
	TransactionCommonFields `xorm:"extends"`

	Owner         string `xorm:"owner varchar(68) null"`
	OfferSequence int64  `xorm:"offer_sequence bigint null"`
}

func (t EscrowCancel) TableName() string {
	return tableName("transaction_escrow_cancel")
}
