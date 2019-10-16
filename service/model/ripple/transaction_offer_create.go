package ripple

type OfferCreate struct {
	TransactionCommonFields `xorm:"extends"`

	Expiration    int64  `xorm:"expiration int null"`
	OfferSequence int64  `xorm:"offer_sequence bigint null"`
	TakerGets     string `xorm:"taker_gets text notnull"`
	TakerPays     string `xorm:"taker_pays text notnull"`
}

func (t OfferCreate) TableName() string {
	return tableName("transaction_offer_create")
}
