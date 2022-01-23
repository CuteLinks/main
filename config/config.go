package config

import (
	"github.com/tkanos/gonfig"
)

type Configuration struct {
	INMEMORYSTORAGE bool
	PORT string
	DB_CONN_STR string
	REDIS_ADDR string
}

func GetConfig(params ...string) Configuration{
	configuration := Configuration{}
	fileName:="config.json"
	gonfig.GetConf(fileName, &configuration)
	return configuration
}