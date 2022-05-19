package repository

import "go-todolist/model"

type TodoRepository interface {
	Insert(todo model.Todo)
	Update(id string, request model.Todo)
	Find(id string) model.Todo
	GetAll() (todos []model.Todo)
	Delete(id string)
}
