package controller

import (
	"go-todolist/model"
	"go-todolist/service"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type TodoController struct {
	TodoService service.TodoService
}

func NewTodoController(todoService *service.TodoService) TodoController {
	return TodoController{TodoService: *todoService}
}

func (controller *TodoController) Create(c *fiber.Ctx) error {
	var request model.Todo
	err := c.BodyParser(&request)
	request.Id = uuid.New().String()
	if err != nil {
		log.Fatal(err.Error())
	}

	response := controller.TodoService.Create(request)
	return c.JSON(fiber.Map{
		"data": response,
	})
}

func (controller *TodoController) GetAll(c *fiber.Ctx) error {
	var todos []model.Todo = controller.TodoService.GetAll()
	return c.JSON(fiber.Map{
		"data": todos,
	})
}

func (controller *TodoController) Update(c *fiber.Ctx) error {
	id := c.Params("id")
	var request model.Todo
	err := c.BodyParser(&request)
	if err != nil {
		log.Fatal(err.Error())
	}

	updatedTodo := controller.TodoService.Update(id, request)

	return c.JSON(fiber.Map{
		"data": updatedTodo,
	})
}

func (controller TodoController) Delete(c *fiber.Ctx) error {
	id := c.Params("id")

	delete := controller.TodoService.Delete(id)

	return c.JSON(fiber.Map{
		"data": delete,
	})
}
