package main

import (
	"fmt"
	"go-todolist/config"

	"github.com/gofiber/fiber/v2"
)

func main() {
	port := fmt.Sprintf(":%s", config.EnvConfig("PORT"))

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World")
	})

	err := app.Listen(port)
	if err != nil {
		panic(err)
	}
}
