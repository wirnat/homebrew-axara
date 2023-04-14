package env

import conf "~import_env~/config"

const (
	ConfigEnv           = "CONFIG_FILE"
	ConfigFile          = "config.yaml"
	MigrationSourcePath = "./infrastructure/migration/source"
	//Message Broker
	MbCustomerTopic = "customer"
	MbMediaTopic    = "media"
	MbCatalogTopic  = "catalog"
	DefaultLogPath  = "./log"
	AppModeDev      = "dev"
	AppModeProd     = "production"
)

var ENV conf.AppConf
