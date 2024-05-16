package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sharantechuser/gotodoservice/controller"
)

func Router() *mux.Router {

	router := mux.NewRouter()

	router.Handle("/api/todos", http.HandlerFunc(controller.GetAllTodos)).Methods("GET")
	router.Handle("/api/todo", http.HandlerFunc(controller.CreateTodo)).Methods("POST")
	router.Handle("/api/todo/{id}", http.HandlerFunc(controller.UpdateTodo)).Methods("PUT")
	router.Handle("/api/todo/{id}", http.HandlerFunc(controller.DeleteTodo)).Methods("DELETE")

	return router
}
