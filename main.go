package main

import (
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
		Prefork:     true,
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(c.App().Stack())
	})

	log.Fatal(app.Listen(":3000"))
}
