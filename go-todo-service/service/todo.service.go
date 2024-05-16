package service

import "github.com/sharantechuser/gotodoservice/model"

type ITodoService interface {
	GetAllTodos() ([]model.Todo, error)

	CreateTodo(todo model.Todo) error
	UpdateTodo(_id string, todo model.Todo) (int64, error)
	DeleteTodo(id string, i interface{}) (int64, error)
}
