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
	if err := database.Migrate(model.Todo{}); err != nil {
		log.Panicln(err)
	}
}
