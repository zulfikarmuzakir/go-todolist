package repository

import "go-todolist/model"

type TodoRepository interface {
	Insert(todo model.Todo)
	GetAll() (todos []model.Todo)
	Delete(param string)
}
