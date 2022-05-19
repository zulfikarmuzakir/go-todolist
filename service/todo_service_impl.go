package service

import (
	"go-todolist/model"
	"go-todolist/repository"
)

type TodoServiceImpl struct {
	TodoRepository repository.TodoRepository
}

func NewTodoService(todoRepository repository.TodoRepository) TodoService {
	return &TodoServiceImpl{}
}

func (service *TodoServiceImpl) Create(request model.Todo) (todo model.Todo) {

}
