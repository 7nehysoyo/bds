package ripple

type SetRegularKey struct {
	TransactionCommonFields `xorm:"extends"`

	RegularKey string `xorm:"regular_key varchar(68) null"`
}

func (t SetRegularKey) TableName() string {
	return tableName("transaction_set_regular_key")
}
