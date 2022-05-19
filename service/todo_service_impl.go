package service

import (
	"go-todolist/model"
	"go-todolist/repository"
)

type TodoServiceImpl struct {
	TodoRepository repository.TodoRepository
}

func NewTodoService(todoRepository *repository.TodoRepository) TodoService {
	return &TodoServiceImpl{
		TodoRepository: *todoRepository,
	}
}

func (service *TodoServiceImpl) Create(request model.Todo) (todo model.Todo) {
	todo = model.Todo{
		Id:          request.Id,
		Name:        request.Name,
		Description: request.Description,
	}

	service.TodoRepository.Insert(todo)

	return todo
}

func (service *TodoServiceImpl) GetAll() (todos []model.Todo) {
	todos = service.TodoRepository.GetAll()
	return todos
}

func (service *TodoServiceImpl) Update(id string, request model.Todo) model.Todo {
	todo := service.TodoRepository.Find(id)
	todo.Name = request.Name
	todo.Description = request.Description

	service.TodoRepository.Update(id, todo)

	updatedTodo := service.TodoRepository.Find(id)

	return updatedTodo
}

func (service *TodoServiceImpl) Delete(id string) string {
	service.TodoRepository.Delete(id)

	return "Success deleted"
}
