package service

import "go-todolist/model"

type TodoService interface {
	Create(request model.Todo) (todo model.Todo)
}
