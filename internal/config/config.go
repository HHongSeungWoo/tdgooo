package config

import (
	"fiber-test/internal/validation"
	"github.com/caarlos0/env/v7"
	"github.com/joho/godotenv"
	"log"
)

type database struct {
	Host     string `env:"HOST" envDefault:"localhost" validate:"required,hostname"`
	Port     int    `env:"PORT" envDefault:"3306" validate:"required,gt=0,lt=65536"`
	User     string `env:"USER,required" validate:"required,gte=1"`
	Password string `env:"PASSWORD"`
	Database string `env:"DATABASE,required" validate:"required,gte=1"`
}

var DB = &database{}

func Init(path ...string) {
	if err := godotenv.Load(path...); err != nil {
		log.Fatal("env 파일을 찾을 수 없습니다.")
	}

	opt := env.Options{
		Prefix: "DB_",
	}
	if err := env.Parse(DB, opt); err != nil {
		log.Fatal(err)
	}
	if err := validation.Struct(DB); err != nil {
		log.Fatal(err)
	}
}
