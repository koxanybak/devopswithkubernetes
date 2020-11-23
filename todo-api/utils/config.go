package utils

import (
	"fmt"
	"os"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"github.com/koxanybak/todo-api/models"
)

// DB is a postgresql database connection
var DB *pg.DB

// ConnectToPostgres allows the use of the variable utils.DB
func ConnectToPostgres() {
	fmt.Println(os.Getenv("POSTGRES_USER"))
	fmt.Println(os.Getenv("POSTGRES_PASSWORD"))
	fmt.Println(os.Getenv("POSTGRES_DB"))

	DB = pg.Connect(&pg.Options{
		Addr: "postgres-svc:5432",
		User: os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		Database: os.Getenv("POSTGRES_DB"),
	})

	err := DB.Model(&models.Todo{}).CreateTable(&orm.CreateTableOptions{
		IfNotExists: true,
	})
	if err != nil {
		panic(err)
	}
}