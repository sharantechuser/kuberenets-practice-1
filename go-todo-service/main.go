package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/sharantechuser/gotodoservice/db"
	"github.com/sharantechuser/gotodoservice/router"
)

func main() {

	if !db.Initialize() {
		log.Fatal("FATAL! could not connect to database")
	}

	fmt.Println("Go with Mongo API")
	r := router.Router()
	fmt.Println("Server is getting started...")
	fmt.Println("Listening at port 4000 ...")

	corsHandler := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
	)

	if err := http.ListenAndServe(":4000", corsHandler(r)); err != nil {
		log.Fatal(err)
	}

}
