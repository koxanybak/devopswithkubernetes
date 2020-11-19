package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/koxanybak/todo-api/models"
	"github.com/koxanybak/todo-api/utils"
)

func getTodos(w http.ResponseWriter, r *http.Request) {
	var todos []models.Todo
	err := utils.DB.Model(&todos).Select()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(todos)
}

func createTodo(w http.ResponseWriter, r *http.Request) {
	var todo models.Todo
	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	newTodo := &models.Todo{
		Name: todo.Name,
	}
	_, err = utils.DB.Model(newTodo).Insert()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newTodo)
}

// TodoRouter adds routes for handling TODO endpoints
func TodoRouter(r *mux.Router) {
	s := r.PathPrefix("/api/todos").Subrouter()

	s.HandleFunc("/", createTodo).Methods("POST")
	s.HandleFunc("/", getTodos).Methods("GET")
}