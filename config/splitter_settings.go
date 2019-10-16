package config

import (
	"fmt"
	"github.com/go-ini/ini"
)

const (
	BCHSection          = "bch"
	BTCSection          = "btc"
	ETCSection          = "etc"
	ETHSection          = "eth"
	EOSSection          = "eos"
	LTCSection          = "ltc"
	DOGESection         = "doge"
	BSVSection          = "bsv"
	TRONSection         = "tron"
	XRPSection          = "xrp"
	XLMSection          = "xlm"
	KafkaBCHSection     = "kafka.bch"
	KafkaBTCSection     = "kafka.btc"
	KafkaETCSection     = "kafka.etc"
	KafkaETHSection     = "kafka.eth"
	KafkaEOSSection     = "kafka.eos"
	KafkaLTCSection     = "kafka.ltc"
	KafkaDOGESection    = "kafka.doge"
	KafkaBSVSection     = "kafka.bsv"
	KafkaTRONSection    = "kafka.tron"
	KafkaXRPSection     = "kafka.xrp"
	KafkaXLMSection     = "kafka.xlm"
	CronBCHSection      = "cron.bch"
	CronBTCSection      = "cron.btc"
	CronETCSection      = "cron.etc"
	CronETHSection      = "cron.eth"
	CronEOSSection      = "cron.eos"
	CronLTCSection      = "cron.ltc"
	CronDOGESection     = "cron.doge"
	CronBSVSection      = "cron.bsv"
	CronTRONSection     = "cron.tron"
	CronXRPSection      = "cron.xrp"
	CronXLMSection      = "cron.xlm"
	DatabaseBTCSection  = "database.btc"
	DatabaseETHSection  = "database.eth"
	DatabaseBCHSection  = "database.bch"
	DatabaseETCSection  = "database.etc"
	DatabaseLTCSection  = "database.ltc"
	DatabaseEOSSection  = "database.eos"
	DatabaseXRPSection  = "database.xrp"
	DatabaseDOGESection = "database.doge"
	DatabaseBSVSection  = "database.bsv"
	DatabaseTRONSection = "database.tron"
	DatabaseXLMSection  = "database.xlm"
)

var SplitterConfig *SplitterSettings

type SplitterSettings struct {
	GlobalSetting       Global
	ProfilingSetting    Profiling
	MetricSetting       Metric
	BCHSetting          BCH
	BTCSetting          BTC
	ETCSetting          ETC
	ETHSetting          ETH
	EOSSetting          EOS
	LTCSetting          LTC
	DOGESetting         DOGE
	BSVSetting          BSV
	TRONSetting         TRON
	XRPSetting          XRP
	XLMSetting          XLM
	KafkaBCHSetting     Kafka
	KafkaBTCSetting     Kafka
	KafkaETCSetting     Kafka
	KafkaETHSetting     Kafka
	KafkaEOSSetting     Kafka
	KafkaTRONSetting    Kafka
	KafkaLTCSetting     Kafka
	KafkaDOGESetting    Kafka
	KafkaBSVSetting     Kafka
	KafkaXRPSetting     Kafka
	KafkaXLMSetting     Kafka
	DatabaseBCHSetting  Database
	DatabaseBTCSetting  Database
	DatabaseETCSetting  Database
	DatabaseETHSetting  Database
	DatabaseEOSSetting  Database
	DatabaseTRONSetting Database
	DatabaseLTCSetting  Database
	DatabaseDOGESetting Database
	DatabaseBSVSetting  Database
	DatabaseXRPSetting  Database
	DatabaseXLMSetting  Database
	CronBCHSetting      CronBCH
	CronBTCSetting      CronBTC
	CronETCSetting      CronETC
	CronETHSetting      CronETH
	CronEOSSetting      CronEOS
	CronLTCSetting      CronLTC
	CronXLMSetting      CronXLM
	CronDOGESetting     CronDOGE
	CronBSVSetting      CronBSV
	CronTRONSetting     CronTRON
	CronXRPSetting      CronXRP
}

type BTC struct {
	Enable                     bool   `ini:"enable"`
	DatabaseEnable             bool   `ini:"database_enable"`
	OmniEnable                 bool   `ini:"omni_enable"`
	MaxBatchBlock              int    `ini:"max_batch_block"`
	Endpoint                   string `ini:"endpoint"`
	User                       string `ini:"user"`
	Password                   string `ini:"password"`
	OmniEndpoint               string `ini:"omni_endpoint"`
	OmniUser                   string `ini:"omni_user"`
	OmniPassword               string `ini:"omni_password"`
	JSONSchemaFile             string `ini:"json_schema_file"`
	JSONSchemaValidationEnable bool   `ini:"json_schema_validation_enable"`
}

type BCH struct {
	Enable                     bool   `ini:"enable"`
	DatabaseEnable             bool   `ini:"database_enable"`
	MaxBatchBlock              int    `ini:"max_batch_block"`
	Endpoint                   string `ini:"endpoint"`
	User                       string `ini:"user"`
	Password                   string `ini:"password"`
	JSONSchemaFile             string `ini:"json_schema_file"`
	JSONSchemaValidationEnable bool   `ini:"json_schema_validation_enable"`
}

type ETC struct {
	Enable                     bool   `ini:"enable"`
	DatabaseEnable             bool   `ini:"database_enable"`
	DatabaseWorkerBuffer       int    `ini:"database_worker_buffer"`
	DatabaseWorkerNumber       int    `ini:"database_worker_number"`
	SkipHeight                 int    `ini:"skip_height"`
	SkipMissBlock              bool   `ini:"skip_miss_block"`
	MaxBatchBlock              int    `ini:"max_batch_block"`
	Endpoint                   string `ini:"endpoint"`
	User                       string `ini:"user"`
	Password                   string `ini:"password"`
	JSONSchemaFile             string `ini:"json_schema_file"`
	JSONSchemaValidationEnable bool   `ini:"json_schema_validation_enable"`
	ConcurrentHeight           int    `ini:"concurrent_height"`
}

type DOGE struct {
	Enable                     bool   `ini:"enable"`
	DatabaseEnable             bool   `ini:"database_enable"`
	MaxBatchBlock              int    `ini:"max_batch_block"`
	Endpoint                   string `ini:"endpoint"`
	User                       string `ini:"user"`
	Password                   string `ini:"password"`
	JSONSchemaFile             string `ini:"json_schema_file"`
	JSONSchemaValidationEnable bool   `ini:"json_schema_validation_enable"`
}

type ETH struct {
	Enable                     bool   `ini:"enable"`
	DatabaseEnable             bool   `ini:"database_enable"`
	DatabaseWorkerBuffer       int    `ini:"database_worker_buffer"`
	DatabaseWorkerNumber       int    `ini:"database_worker_number"`
	MaxBatchBlock              int    `ini:"max_batch_block"`
	Endpoint                   string `ini:"endpoint"`
	User                       string `ini:"user"`
	Password                   string `ini:"password"`
	JSONSchemaFile             string `ini:"json_schema_file"`
	JSONSchemaValidationEnable bool   `ini:"json_schema_validation_enable"`
	ConcurrentHeight           int    `ini:"concurrent_height"`
}

type LTC struct {
	Enable                     bool   `ini:"enable"`
	DatabaseEnable             bool   `ini:"database_enable"`
	MaxBatchBlock              int    `ini:"max_batch_block"`
	Endpoint                   string `ini:"endpoint"`
	User                       string `ini:"user"`
	Password                   string `ini:"password"`
	JSONSchemaFile             string `ini:"json_schema_file"`
	JSONSchemaValidationEnable bool   `ini:"json_schema_validation_enable"`
}

type EOS struct {
	Enable                     bool   `ini:"enable"`
	DatabaseEnable             bool   `ini:"database_enable"`
	DatabaseWorkerBuffer       int    `ini:"database_worker_buffer"`
	DatabaseWorkerNumber       int    `ini:"database_worker_number"`
	SkipHeight                 int    `ini:"skip_height"`
	SkipMissBlock              bool   `ini:"skip_miss_block"`
	MaxBatchBlock              int    `ini:"max_batch_block"`
	Endpoint                   string `ini:"endpoint"`
	KafkaProxyHost             string `ini:"kafka_proxy_host"`
	KafkaProxyPort             string `ini:"kafka_proxy_port"`
	User                       string `ini:"user"`
	Password                   string `ini:"password"`
	JSONSchemaFile             string `ini:"json_schema_file"`
	JSONSchemaValidationEnable bool   `ini:"json_schema_validation_enable"`
}

type BSV struct {
	Enable                     bool   `ini:"enable"`
	DatabaseEnable             bool   `ini:"database_enable"`
	MaxBatchBlock              int    `ini:"max_batch_block"`
	Endpoint                   string `ini:"endpoint"`
	User                       string `ini:"user"`
	Password                   string `ini:"password"`
	JSONSchemaFile             string `ini:"json_schema_file"`
	JSONSchemaValidationEnable bool   `ini:"json_schema_validation_enable"`
}

type TRON struct {
	Enable                     bool   `ini:"enable"`
	DatabaseEnable             bool   `ini:"database_enable"`
	DatabaseWorkerBuffer       int    `ini:"database_worker_buffer"`
	DatabaseWorkerNumber       int    `ini:"database_worker_number"`
	SkipHeight                 int    `ini:"skip_height"`
	SkipMissBlock              bool   `ini:"skip_miss_block"`
	MaxBatchBlock              int    `ini:"max_batch_block"`
	ConcurrentHeight           int    `ini:"concurrent_height"`
	Endpoint                   string `ini:"endpoint"`
	User                       string `ini:"user"`
	Password                   string `ini:"password"`
	JSONSchemaFile             string `ini:"json_schema_file"`
	JSONSchemaValidationEnable bool   `ini:"json_schema_validation_enable"`
}

type XRP struct {
	Enable                     bool   `ini:"enable"`
	DatabaseEnable             bool   `ini:"database_enable"`
	DatabaseWorkerBuffer       int    `ini:"database_worker_buffer"`
	DatabaseWorkerNumber       int    `ini:"database_worker_number"`
	MaxBatchBlock              int    `ini:"max_batch_block"`
	Endpoint                   string `ini:"endpoint"`
	KafkaProxyHost             string `ini:"kafka_proxy_host"`
	KafkaProxyPort             string `ini:"kafka_proxy_port"`
	User                       string `ini:"user"`
	Password                   string `ini:"password"`
	JSONSchemaFile             string `ini:"json_schema_file"`
	JSONSchemaValidationEnable bool   `ini:"json_schema_validation_enable"`
}

type XLM struct {
	Enable                     bool   `ini:"enable"`
	DatabaseEnable             bool   `ini:"database_enable"`
	DatabaseWorkerBuffer       int    `ini:"database_worker_buffer"`
	DatabaseWorkerNumber       int    `ini:"database_worker_number"`
	MaxBatchBlock              int    `ini:"max_batch_block"`
	Endpoint                   string `ini:"endpoint"`
	JSONSchemaFile             string `ini:"json_schema_file"`
	JSONSchemaValidationEnable bool   `ini:"json_schema_validation_enable"`
	ConcurrentHeight           int64  `ini:"concurrent_height"`
	ConcurrentHTTP             int64  `ini:"concurrent_http"`
}

type CronBCH struct {
	UpdateMetaExpr    string `ini:"update_meta_expr"`
	GetBatchBlockExpr string `ini:"get_batch_block_expr"`
}

type CronBTC struct {
	UpdateMetaExpr    string `ini:"update_meta_expr"`
	GetBatchBlockExpr string `ini:"get_batch_block_expr"`
}

type CronETC struct {
	UpdateMetaExpr             string `ini:"update_meta_expr"`
	GetBatchBlockExpr          string `ini:"get_batch_block_expr"`
	RefreshContractAddressExpr string `ini:"refresh_contract_address_expr"`
	RefreshPoolNameExpr        string `ini:"refresh_pool_name_expr"`
}

type CronETH struct {
	UpdateMetaExpr             string `ini:"update_meta_expr"`
	GetBatchBlockExpr          string `ini:"get_batch_block_expr"`
	RefreshContractAddressExpr string `ini:"refresh_contract_address_expr"`
	RefreshPoolNameExpr        string `ini:"refresh_pool_name_expr"`
}

type CronLTC struct {
	UpdateMetaExpr    string `ini:"update_meta_expr"`
	GetBatchBlockExpr string `ini:"get_batch_block_expr"`
}

type CronEOS struct {
	UpdateMetaExpr string `ini:"update_meta_expr"`
}

type CronDOGE struct {
	UpdateMetaExpr    string `ini:"update_meta_expr"`
	GetBatchBlockExpr string `ini:"get_batch_block_expr"`
}

type CronBSV struct {
	UpdateMetaExpr    string `ini:"update_meta_expr"`
	GetBatchBlockExpr string `ini:"get_batch_block_expr"`
}

type CronTRON struct {
	UpdateMetaExpr string `ini:"update_meta_expr"`
}
type CronXRP struct {
	UpdateMetaExpr     string `ini:"update_meta_expr"`
	GetBatchLedgerExpr string `ini:"get_batch_ledger_expr"`
}

type CronXLM struct {
	UpdateMetaExpr string `ini:"update_meta_expr"`
}

type Kafka struct {
	Topic        string
	ClientID     string `ini:"client_id"`
	GroupID      string `ini:"group_id"`
	BrokerList   string
	BufferSize   int
	ReturnErrors bool `ini:"return_errors"`
}

func InitSplitterConfig(config string) (err error) {
	var cfg *ini.File
	cfg, err = ini.Load(config)
	if err != nil {
		fmt.Println("Read config file error: " + config)
		return err
	}
	cfg.NameMapper = ini.TitleUnderscore

	SplitterConfig = new(SplitterSettings)
	err = cfg.Section(GlobalSection).MapTo(&SplitterConfig.GlobalSetting)
	if err != nil {
		return err
	}
	err = cfg.Section(ProfilingSection).MapTo(&SplitterConfig.ProfilingSetting)
	if err != nil {
		return err
	}
	err = cfg.Section(MetricSection).MapTo(&SplitterConfig.MetricSetting)
	if err != nil {
		return err
	}

	err = cfg.Section(DatabaseBCHSection).MapTo(&SplitterConfig.DatabaseBCHSetting)
	if err != nil {
		return err
	}
	err = cfg.Section(DatabaseBTCSection).MapTo(&SplitterConfig.DatabaseBTCSetting)
	if err != nil {
		return err
	}
	err = cfg.Section(DatabaseETCSection).MapTo(&SplitterConfig.DatabaseETCSetting)
	if err != nil {
		return err
	}
	err = cfg.Section(DatabaseETHSection).MapTo(&SplitterConfig.DatabaseETHSetting)
	if err != nil {
		return err
	}
	err = cfg.Section(DatabaseEOSSection).MapTo(&SplitterConfig.DatabaseEOSSetting)
	if err != nil {
		return err
	}
	err = cfg.Section(DatabaseLTCSection).MapTo(&SplitterConfig.DatabaseLTCSetting)
	if err != nil {
		return err
	}
	err = cfg.Section(DatabaseDOGESection).MapTo(&SplitterConfig.DatabaseDOGESetting)
	if err != nil {
		return err
	}
	err = cfg.Section(DatabaseBSVSection).MapTo(&SplitterConfig.DatabaseBSVSetting)
	if err != nil {
		return err
	}
	err = cfg.Section(DatabaseTRONSection).MapTo(&SplitterConfig.DatabaseTRONSetting)
	if err != nil {
		return err
	}
	err = cfg.Section(DatabaseXRPSection).MapTo(&SplitterConfig.DatabaseXRPSetting)
	if err != nil {
		return err
	}
	err = cfg.Section(DatabaseXLMSection).MapTo(&SplitterConfig.DatabaseXLMSetting)
	if err != nil {
		return err
	}

	err = cfg.Section(KafkaBCHSection).MapTo(&SplitterConfig.KafkaBCHSetting)
	if err != nil {
		return err
	}
	err = cfg.Section(KafkaBTCSection).MapTo(&SplitterConfig.KafkaBTCSetting)
	if err != nil {
		return err
	}
	err = cfg.Section(KafkaETCSection).MapTo(&SplitterConfig.KafkaETCSetting)
	if err != nil {
		return err
	}
	err = cfg.Section(KafkaETHSection).MapTo(&SplitterConfig.KafkaETHSetting)
	if err != nil {
		return err
	}
	err = cfg.Section(KafkaEOSSection).MapTo(&SplitterConfig.KafkaEOSSetting)
	if err != nil {
		return err
	}
	err = cfg.Section(KafkaLTCSection).MapTo(&SplitterConfig.KafkaLTCSetting)
	if err != nil {
		return err
	}
	err = cfg.Section(KafkaDOGESection).MapTo(&SplitterConfig.KafkaDOGESetting)
	if err != nil {
		return err
	}
	err = cfg.Section(KafkaBSVSection).MapTo(&SplitterConfig.KafkaBSVSetting)
	if err != nil {
		return err
	}
	err = cfg.Section(KafkaTRONSection).MapTo(&SplitterConfig.KafkaTRONSetting)
	if err != nil {
		return err
	}
	err = cfg.Section(KafkaXRPSection).MapTo(&SplitterConfig.KafkaXRPSetting)
	if err != nil {
		return err
	}
	err = cfg.Section(KafkaXLMSection).MapTo(&SplitterConfig.KafkaXLMSetting)
	if err != nil {
		return err
	}

	err = cfg.Section(CronBCHSection).MapTo(&SplitterConfig.CronBCHSetting)
	if err != nil {
		return err
	}
	err = cfg.Section(CronBTCSection).MapTo(&SplitterConfig.CronBTCSetting)
	if err != nil {
		return err
	}
	err = cfg.Section(CronETCSection).MapTo(&SplitterConfig.CronETCSetting)
	if err != nil {
		return err
	}
	err = cfg.Section(CronETHSection).MapTo(&SplitterConfig.CronETHSetting)
	if err != nil {
		return err
	}
	err = cfg.Section(CronEOSSection).MapTo(&SplitterConfig.CronEOSSetting)
	if err != nil {
		return err
	}
	err = cfg.Section(CronLTCSection).MapTo(&SplitterConfig.CronLTCSetting)
	if err != nil {
		return err
	}
	err = cfg.Section(CronDOGESection).MapTo(&SplitterConfig.CronDOGESetting)
	if err != nil {
		return err
	}
	err = cfg.Section(CronBSVSection).MapTo(&SplitterConfig.CronBSVSetting)
	if err != nil {
		return err
	}
	err = cfg.Section(CronTRONSection).MapTo(&SplitterConfig.CronTRONSetting)
	if err != nil {
		return err
	}
	err = cfg.Section(CronXRPSection).MapTo(&SplitterConfig.CronXRPSetting)
	if err != nil {
		return err
	}
	err = cfg.Section(CronXLMSection).MapTo(&SplitterConfig.CronXLMSetting)
	if err != nil {
		return err
	}

	err = cfg.Section(BCHSection).MapTo(&SplitterConfig.BCHSetting)
	if err != nil {
		return err
	}
	err = cfg.Section(BTCSection).MapTo(&SplitterConfig.BTCSetting)
	if err != nil {
		return err
	}
	err = cfg.Section(ETCSection).MapTo(&SplitterConfig.ETCSetting)
	if err != nil {
		return err
	}
	err = cfg.Section(ETHSection).MapTo(&SplitterConfig.ETHSetting)
	if err != nil {
		return err
	}
	err = cfg.Section(EOSSection).MapTo(&SplitterConfig.EOSSetting)
	if err != nil {
		return err
	}
	err = cfg.Section(LTCSection).MapTo(&SplitterConfig.LTCSetting)
	if err != nil {
		return err
	}
	err = cfg.Section(DOGESection).MapTo(&SplitterConfig.DOGESetting)
	if err != nil {
		return err
	}
	err = cfg.Section(BSVSection).MapTo(&SplitterConfig.BSVSetting)
	if err != nil {
		return err
	}
	err = cfg.Section(TRONSection).MapTo(&SplitterConfig.TRONSetting)
	if err != nil {
		return err
	}
	err = cfg.Section(XRPSection).MapTo(&SplitterConfig.XRPSetting)
	if err != nil {
		return err
	}
	err = cfg.Section(XLMSection).MapTo(&SplitterConfig.XLMSetting)
	if err != nil {
		return err
	}
	return nil
}
