package database

import (
	"errors"
	"fiber-test/internal/config"
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"os"
	"runtime"
)

var DB *gorm.DB

func Connect() (err error) {
	if DB != nil {
		// TODO warn level 의 로그가 남도록 수정
		log.Println("warn 이미 연결된 Connection 이 있습니다.")
		return
	}

	log.Println(runtime.Caller(0))
	log.Println(os.Executable())
	log.Println(os.Getwd())

	DB, err = gorm.Open(sqlite.Open(fmt.Sprintf("%s.db", config.DB.Database)), &gorm.Config{})
	return
}

func MustConnect() {
	if err := Connect(); err != nil {
		log.Fatal(err)
	}
}

func Migrate(models ...interface{}) (err error) {
	if DB == nil {
		return errors.New("`database.Connect()`을 먼저 실행해주세요")
	}
	err = DB.AutoMigrate(models...)
	return
}

func Close() {
	if DB == nil {
		return
	}

	if db, err := DB.DB(); err != nil {
		_ = db.Close()
	}
	DB = nil
}
