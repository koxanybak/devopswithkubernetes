package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/koxanybak/todo-api/handlers"
	"github.com/koxanybak/todo-api/utils"
)

func main() {
	utils.ConnectToPostgres()

	r := mux.NewRouter()
	handlers.TodoRouter(r)

	log.Println("Listening on 8000")
	log.Fatal(http.ListenAndServe("0.0.0.0:8000", r))
}