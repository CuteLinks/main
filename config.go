package main

import (
	"github.com/tkanos/gonfig"
)

type Configuration struct {
	DB_USERNAME string
	INMEMORYSTORAGE bool
}

func GetConfig(params ...string) Configuration{
	configuration := Configuration{}
	fileName:="config.json"
	gonfig.GetConf(fileName, &configuration)
	return configuration
}