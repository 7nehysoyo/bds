package ripple

type Payment struct {
	TransactionCommonFields `xorm:"extends"`

	Amount         string `xorm:"amount text notnull"`
	Destination    string `xorm:"destination varchar(68) notnull"`
	DestinationTag int64  `xorm:"destination_tag bigint null"`
	InvoiceID      string `xorm:"invoice_id varchar(128) null"`
	Paths          string `xorm:"paths text null"`
	SendMax        string `xorm:"send_max text null"`
	DeliverMin     string `xorm:"deliver_min text null"`
}

func (t Payment) TableName() string {
	return tableName("transaction_payment")
}
