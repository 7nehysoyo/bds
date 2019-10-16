package ripple

type PaymentChannelFund struct {
	TransactionCommonFields `xorm:"extends"`

	Channel    string `xorm:"channel char(64) notnull"`
	Amount     string `xorm:"amount text notnull"`
	Expiration int64  `xorm:"expiration int null"`
}

func (t PaymentChannelFund) TableName() string {
	return tableName("transaction_payment_channel_fund")
}
