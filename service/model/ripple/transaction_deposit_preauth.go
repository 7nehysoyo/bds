package ripple

type DepositPreauth struct {
	TransactionCommonFields `xorm:"extends"`

	Authorize   string `xorm:"authorize varchar(68) null"`
	UnAuthorize string `xorm:"un_authorize varchar(68) null"`
}

func (t DepositPreauth) TableName() string {
	return tableName("transaction_deposit_preauth")
}
