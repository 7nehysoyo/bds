package ripple

type PaymentChannelClaim struct {
	TransactionCommonFields `xorm:"extends"`

	Channel   string `xorm:"channel char(64) notnull"`
	Amount    string `xorm:"amount text null"`
	Balance   string `xorm:"balance text null"`
	PublicKey string `xorm:"public_key char(66) null"`
	Signature string `xorm:"signature text null"`
}

func (t PaymentChannelClaim) TableName() string {
	return tableName("transaction_payment_channel_claim")
}
