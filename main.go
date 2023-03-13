package main

import (
	"fiber-test/internal/config"
	"fiber-test/internal/database"
	"fiber-test/internal/model"
	"fiber-test/internal/validation"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"log"
)

type CreateTodo struct {
	Title       string  `json:"title" validate:"required,min=2"`
	Category    *string `json:"category,omitempty"`
	Description *string `json:"description,omitempty"`
}

type UpdateTodo struct {
	Title       *string `json:"title,omitempty" validate:"omitempty,min=2"`
	Category    *string `json:"category,omitempty"`
	Description *string `json:"description,omitempty"`
}

type Pagination struct {
	Limit  int `query:"limit"`
	Offset int `query:"offset"`
}

func main() {
	config.MustInit(".env")
	database.MustConnect()

	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	app.Get("/", func(c *fiber.Ctx) error {
		todos := &[]model.Todo{}
		query := &Pagination{}
		if err := c.QueryParser(query); err != nil {
			c.Status(400)
			return nil
		}

		if query.Limit == 0 {
			query.Limit = 10
		}

		database.DB.Order("id desc").Limit(query.Limit).Offset(query.Offset).Find(&todos)
		return c.JSON(todos)
	})

	app.Post("/", func(c *fiber.Ctx) error {
		body := &CreateTodo{}

		if err := c.BodyParser(body); err != nil {
			return err
		}

		if err := validation.Struct(body); err != nil {
			return err
		}

		todo := &model.Todo{
			Title: body.Title,
		}
		if body.Category != nil {
			todo.Category = model.NullString{
				String: *body.Category,
				Valid:  true,
			}
		}
		if body.Description != nil {
			todo.Description = model.NullString{
				String: *body.Description,
				Valid:  true,
			}
		}
		database.DB.Save(todo)

		return c.JSON(&todo)
	})

	app.Get("/:id", func(c *fiber.Ctx) error {
		id, _ := c.ParamsInt("id", -1)
		if id <= 0 {
			c.Status(404)
			return nil
		}
		todo := &model.Todo{Id: uint(id)}
		if database.DB.Find(todo).RowsAffected != 1 {
			c.Status(404)
			return nil
		}

		return c.JSON(todo)
	})

	app.Patch("/:id", func(c *fiber.Ctx) error {
		id, _ := c.ParamsInt("id", -1)
		if id <= 0 {
			c.Status(404)
			return nil
		}

		body := &UpdateTodo{}

		if err := c.BodyParser(body); err != nil {
			return err
		}

		if err := validation.Struct(body); err != nil {
			return err
		}

		todo := &model.Todo{Id: uint(id)}
		if database.DB.Find(todo).RowsAffected != 1 {
			c.Status(404)
			return nil
		}

		if body.Title != nil {
			todo.Title = *body.Title
		}
		if body.Category != nil {
			todo.Category = model.NullString{
				String: *body.Category,
				Valid:  true,
			}
		}
		if body.Description != nil {
			todo.Description = model.NullString{
				String: *body.Description,
				Valid:  true,
			}
		}

		database.DB.Save(todo)
		return nil
	})

	app.Delete("/:id", func(c *fiber.Ctx) error {
		id, _ := c.ParamsInt("id", -1)
		if id <= 0 {
			c.Status(404)
			return nil
		}
		todo := &model.Todo{Id: uint(id)}
		database.DB.Delete(todo)

		return nil
	})

	log.Fatal(app.Listen(":3000"))
}
