package ripple

type TrustSet struct {
	TransactionCommonFields `xorm:"extends"`

	LimitAmount string `xorm:"limit_amount text notnull"`
	QualityIn   int64  `xorm:"quality_in bigint null"`
	QualityOut  int64  `xorm:"quality_out bigint null"`
}

func (t TrustSet) TableName() string {
	return tableName("transaction_trust_set")
}
