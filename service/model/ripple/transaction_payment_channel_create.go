package ripple

type PaymentChannelCreate struct {
	TransactionCommonFields `xorm:"extends"`

	Amount         string `xorm:"amount text notnull"`
	Destination    string `xorm:"destination varchar(68) notnull"`
	SettleDelay    int64  `xorm:"settleDelay bigint notnull"`
	PublicKey      string `xorm:"public_key char(66) notnull"`
	CancelAfter    int64  `xorm:"cancel_after bigint null"`
	DestinationTag int64  `xorm:"destination_tag bigint null"`
}

func (t PaymentChannelCreate) TableName() string {
	return tableName("transaction_payment_channel_create")
}
