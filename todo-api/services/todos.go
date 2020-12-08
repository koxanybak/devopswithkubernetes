package services

import (
	"log"

	"github.com/koxanybak/todo-api/models"
	"github.com/koxanybak/todo-api/utils"
)

// GetTodos gets all todos
func GetTodos() []models.Todo {
	var todos []models.Todo
	err := utils.DB.Model(&todos).Select()
	if err != nil {
		panic(err)
	}

	return todos
}

// CreateTodo creates a todo
func CreateTodo(name string) *models.Todo {
	newTodo := &models.Todo{
		Name: name,
		Done: false,
	}
	_, err := utils.DB.Model(newTodo).Insert()
	if err != nil {
		panic(err)
	}
	log.Println("New todo created:", newTodo)
	return newTodo
}

// UpdateTodo updates a todo
func UpdateTodo(newTodo *models.Todo) *models.Todo {
	_, err := utils.DB.Model(newTodo).
		OnConflict("(id) DO UPDATE").
		Set("Name = EXCLUDED.Name, Done = EXCLUDED.Done").
		Insert()
	if err != nil {
		panic(err)
	}
	log.Println("Todo updated:", newTodo)
	return newTodo
}