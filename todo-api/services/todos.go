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

// CreateTodo gets all todos
func CreateTodo(name string) *models.Todo {
	newTodo := &models.Todo{
		Name: name,
	}
	_, err := utils.DB.Model(newTodo).Insert()
	if err != nil {
		panic(err)
	}
	log.Println("New todo created:", newTodo)
	return newTodo
}