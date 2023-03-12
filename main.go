package main

import (
	"fiber-test/internal/config"
	"fiber-test/internal/database"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	config.MustInit(".env")
	database.MustConnect()

	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return nil
	})

	app.Post("/", func(c *fiber.Ctx) error {
		return nil
	})

	app.Get("/:id", func(c *fiber.Ctx) error {
		return nil
	})

	app.Patch("/:id", func(c *fiber.Ctx) error {
		return nil
	})

	app.Delete("/:id", func(c *fiber.Ctx) error {
		return nil
	})

	log.Fatal(app.Listen(":3000"))
}
