package service

import (
	"go-todolist/model"
)

type TodoService interface {
	Create(request model.Todo) (todo model.Todo)
	Update(id string, request model.Todo) model.Todo
	GetAll() (todos []model.Todo)
	Delete(id string) string
}
