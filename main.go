package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"go-todolist/config"
)

func main() {
	port := fmt.Sprintf(":%s", config.EnvConfig("PORT"))
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World")
	})

	app.Listen(port)
}
