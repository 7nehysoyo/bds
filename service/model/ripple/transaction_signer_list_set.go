package ripple

type SignerListSet struct {
	TransactionCommonFields `xorm:"extends"`

	SignerQuorum  int64  `xorm:"signer_quorum bigint notnull"`
	SignerEntries string `xorm:"signer_entries text null"`
}

func (t SignerListSet) TableName() string {
	return tableName("transaction_set_regular_key")
}
