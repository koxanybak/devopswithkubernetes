package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/koxanybak/todo-api/models"
	"github.com/koxanybak/todo-api/utils"
)

func sendToNats(todo *models.Todo) {
	msgData, err := json.Marshal(todo)
	if err != nil {
		log.Println("Couldn't marshal json:", err.Error())
	}
	err = utils.NC.Publish("todo", msgData)
	if err != nil {
		log.Println("Couldn't send msg to NATS:", err.Error())
		utils.NC.Status()
	}
}

func getTodos(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL.Path)
	var todos []models.Todo
	err := utils.DB.Model(&todos).Select()
	if err != nil {
		log.Println(err)
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
		Done: false,
	}

	log.Println("New TODO haloo:", newTodo)
	// Check length
	if len(newTodo.Name) > 140 {
		http.Error(w, "Too long name for TODO", http.StatusBadRequest)
		return
	}

	// Insert
	_, err = utils.DB.Model(newTodo).Insert()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	sendToNats(newTodo)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newTodo)
}

func updateTodo(w http.ResponseWriter, r *http.Request) {
	var todo models.Todo
	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	newTodo := &todo

	log.Println("New TODO:", newTodo)
	// Check length
	if len(newTodo.Name) > 140 {
		http.Error(w, "Too long name for TODO", http.StatusBadRequest)
		return
	}

	if (newTodo.ID <= 0) {
		http.Error(w, "Missing or wrong id", http.StatusBadRequest)
		return
	}

	// Insert
	_, err = utils.DB.Model(newTodo).
		Column("name", "done").
		WherePK().
		Update()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	sendToNats(newTodo)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newTodo)
}

// TodoRouter adds routes for handling TODO endpoints
func TodoRouter(r *mux.Router) {
	s := r.PathPrefix("/api/todos").Subrouter()

	s.HandleFunc("/", createTodo).Methods("POST")
	s.HandleFunc("/", updateTodo).Methods("PUT")
	s.HandleFunc("/", getTodos).Methods("GET")
}