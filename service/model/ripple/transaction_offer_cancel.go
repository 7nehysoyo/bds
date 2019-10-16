package ripple

type OfferCancel struct {
	TransactionCommonFields `xorm:"extends"`

	OfferSequence int64 `xorm:"offer_sequence bigint notnull"`
}

func (t OfferCancel) TableName() string {
	return tableName("transaction_offer_cancel")
}
