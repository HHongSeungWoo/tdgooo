package main

import (
	"fiber-test/internal/config"
	"fiber-test/internal/database"
	"fiber-test/internal/model"
	"log"
)

func main() {
	config.MustInit(".env")
	database.MustConnect()
	if err := database.Migrate(model.TODO{}); err != nil {
		log.Panicln(err)
	}
}
