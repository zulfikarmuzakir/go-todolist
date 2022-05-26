package main

import (
	"fmt"
	"go-todolist/config"
	"go-todolist/controller"
	"go-todolist/repository"
	"go-todolist/service"

	"github.com/gofiber/fiber/v2"
)

func main() {
	port := fmt.Sprintf(":%s", config.EnvConfig("PORT"))

	database := config.ConnectDB()
	todoRepository := repository.NewTodoRepository(database)
	todoService := service.NewTodoService(&todoRepository)
	todoController := controller.NewTodoController(&todoService)

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World")
	})

	app.Get("/todos", todoController.GetAll)
	app.Post("/todos", todoController.Create)
	app.Put("/todos/:id", todoController.Update)
	app.Delete("/todos/:id", todoController.Delete)

	err := app.Listen(port)
	if err != nil {
		panic(err)
	}
}
