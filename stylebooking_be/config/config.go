package config

import (
	"fmt"

	"github.com/tkanos/gonfig"
)

type Configuration struct {
	App   AppConfiguration
	Mongo MongoConfiguration
	Jwt   JwtConfiguration
}

type AppConfiguration struct {
	Addr string
}

type MongoConfiguration struct {
	ConnectionString string
}

type JwtConfiguration struct {
	Secret               string
	Issuer               string
	Audience             string
	AccessTokenDuration  int
	RefreshTokenDuration int
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
