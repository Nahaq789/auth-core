package conf

import (
	"log"

	"github.com/caarlos0/env/v11"
)

type appSetting struct {
	Port string `env:"PORT"`
	Addr string `env:"ADDR"`
}

func SetEnv() appSetting {
	var conf appSetting
	if err := env.Parse(&conf); err != nil {
		log.Fatal(err)
	}
	return conf
}
