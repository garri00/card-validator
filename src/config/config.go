package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Configs struct {
	Host     string `envconfig:"HOST" default:"localhost"`
	Port     string `envconfig:"PORT" default:"8080"`
	LogLevel string `envconfig:"LOGLEVEL"`
}

func GetConfig() (Configs, error) {
	var cfg Configs

	err := godotenv.Load()
	if err != nil {
		return cfg, err
	}

	err = envconfig.Process("myApp", &cfg)
	if err != nil {
		return cfg, err
	}

	return cfg, nil
}
