package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/koxanybak/todo-api/handlers"
	"github.com/koxanybak/todo-api/utils"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func main() {
	utils.ConnectToPostgres()
	utils.NatsConnect()

	r := mux.NewRouter()
	r.HandleFunc("/", home)
	handlers.TodoRouter(r)

	log.Println("Listening on 8000")
	log.Fatal(http.ListenAndServe("0.0.0.0:8000", r))
}