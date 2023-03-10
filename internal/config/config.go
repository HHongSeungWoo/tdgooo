package config

import (
	"github.com/caarlos0/env/v7"
	"github.com/joho/godotenv"
	"log"
)

type DatabaseConfig struct {
	Host     string `env:"HOST" envDefault:"localhost"`
	Port     int    `env:"PORT" envDefault:"3306"`
	User     string `env:"USER,required"`
	Password string `env:"PASSWORD,required"`
	Database string `env:"DATABASE,required"`
}

var DBConfig = &DatabaseConfig{}

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	opt := env.Options{
		Prefix: "DB_",
	}
	if err := env.Parse(DBConfig, opt); err != nil {
		log.Fatal(err)
	}
}
