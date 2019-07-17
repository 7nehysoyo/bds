package etc

import (
	"github.com/jdcloud-bds/bds/common/math"
)

type Account struct {
	ID                      int64                `xorm:"id bigint autoincr pk"`
	Address                 string               `xorm:"address char(40) notnull unique index"`             //账户地址
	Type                    int                  `xorm:"type tinyint notnull default '0'"`                  //账户类型
	DeadAccount             int                  `xorm:"dead_account smallint notnull default '0' "`        //死账户
	Balance                 math.HexOrDecimal256 `xorm:"balance decimal(30,0) notnull default '0'"`         //账户余额
	Creator                 string               `xorm:"creator char(40) notnull default ''"`               //合约部署者
	ContractCreateTimestamp int64                `xorm:"contract_create_timestamp int notnull default '0'"` //合约创建时间
	MinerCount              int                  `xorm:"miner_count int notnull default '0'"`               //矿工挖到的块数
	MinerUncleCount         int                  `xorm:"miner_uncle_count int notnull default '0'"`         //矿工挖到的叔块数量
	BirthTimestamp          int64                `xorm:"birth_timestamp int notnull default '0' "`          //账户第一次出现的时间
	LastActiveTimestamp     int64                `xorm:"last_active_timestamp int notnull default '0'"`     //账户上次活跃时间
}

func (t Account) TableName() string {
	return tableName("account")
}
