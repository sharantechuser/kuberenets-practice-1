package controller

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sharantechuser/gotodoservice/model"
	"github.com/sharantechuser/gotodoservice/service"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type TodoController struct {
	TodoService service.ITodoService
}

func New(todoservice service.ITodoService) TodoController {

	return TodoController{
		TodoService: todoservice,
	}

}

const (
	connectionString = "localhost:27917"
	db_name          = "tododb"
	coll_name        = "todo"
)

var Todo_collection *mongo.Collection

func init() {
	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Disconnect the client when done
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			log.Fatal(err)
		}
	}()

	// Access the "test" database and the "people" collection
	Todo_collection = client.Database(db_name).Collection(coll_name)
}

func GetAllTodos(w http.ResponseWriter, r *http.Request) {

	log.Printf("call GetAllTodos")
	todos, err := service.NewTodoService().GetAllTodos()
	if err != nil {
		log.Printf("ERROR! %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}

	bytes, err := json.Marshal(todos)
	if err != nil {
		log.Printf("ERROR! %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.Write(bytes)
	w.WriteHeader(http.StatusCreated)
}

func CreateTodo(w http.ResponseWriter, r *http.Request) {

	payload, _ := io.ReadAll(r.Body)

	var todo model.Todo
	json.Unmarshal(payload, &todo)

	err := service.NewTodoService().CreateTodo(todo)
	if err != nil {
		log.Printf("ERROR! %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusCreated)

}
func UpdateTodo(w http.ResponseWriter, r *http.Request) {

	var todo model.Todo

	params := mux.Vars(r)
	_id := params["id"]

	payload, _ := io.ReadAll(r.Body)
	json.Unmarshal(payload, &todo)

	_, err := service.NewTodoService().UpdateTodo(_id, todo)
	if err != nil {
		log.Printf("ERROR! %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
}
func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	_id := params["id"]

	_, err := service.NewTodoService().DeleteTodo(_id, model.Todo{})
	if err != nil {
		log.Printf("ERROR! %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
}
