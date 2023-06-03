package config

import (
	"fmt"

	"github.com/tkanos/gonfig"
)

type Configuration struct {
	Mongo MongoConfiguration
}

type MongoConfiguration struct {
	ConnectionString string
}

func GetConfig(params ...string) Configuration {
	configuration := Configuration{}
	env := "dev"
	if len(params) > 0 {
		env = params[0]
	}
	fileName := fmt.Sprintf("../%s_config.json", env)
	err := gonfig.GetConf(fileName, &configuration)
	if err != nil {
		panic(err)
	}
	return configuration
}
