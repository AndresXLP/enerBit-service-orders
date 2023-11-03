package config

import (
	"log"
	"sync"

	"github.com/andresxlp/gosuite/config"
)

type Config struct {
	GRPC     GRPC     `mapstructure:"grpc" validate:"required"`
	Postgres Postgres `mapstructure:"postgres" validate:"required"`
	Redis    Redis    `mapstructure:"redis" validate:"required"`
}

type GRPC struct {
	Protocol string `mapstructure:"protocol" validate:"required"`
	Host     string `mapstructure:"host" validate:"required"`
	Port     int    `mapstructure:"port" validate:"required"`
}

type Postgres struct {
	Host   string `mapstructure:"host" validate:"required"`
	Port   int    `mapstructure:"port" validate:"required"`
	User   string `mapstructure:"user" validate:"required"`
	Pass   string `mapstructure:"pass" validate:"required"`
	DbName string `mapstructure:"db_name" validate:"required"`
}

type Redis struct {
	Host string `mapstructure:"host" validate:"required"`
	Port int    `mapstructure:"port" validate:"required"`
}

var (
	once sync.Once
	Cfg  Config
)

func Environments() Config {
	once.Do(func() {
		if err := config.GetConfigFromEnv(&Cfg); err != nil {
			log.Panicf("Error parsing environments vars %v\n", err)
		}
	})

	return Cfg
}
