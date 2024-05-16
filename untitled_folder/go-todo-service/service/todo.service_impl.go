package service

import (
	"github.com/sharantechuser/gotodoservice/db"
	"github.com/sharantechuser/gotodoservice/model"
)

const (
	todo_coll_name = "Todo"
)

type TodoServiceImpl struct {
}

func NewTodoService() ITodoService {
	return &TodoServiceImpl{}
}

func (us *TodoServiceImpl) GetAllTodos() ([]model.Todo, error) {

	var results []model.Todo
	result, err := db.Handle.GetAll(todo_coll_name, results)

	if err != nil {
		return []model.Todo{}, err
	}

	return result.([]model.Todo), err
}

func (us *TodoServiceImpl) CreateTodo(todo model.Todo) error {
	return db.Handle.Create(todo)
}
func (us *TodoServiceImpl) UpdateTodo(_id string, todo model.Todo) (int64, error) {
	return db.Handle.Update(_id, todo)
}
func (us *TodoServiceImpl) DeleteTodo(_id string, todo interface{}) (int64, error) {
	return db.Handle.Update(_id, todo)
}
