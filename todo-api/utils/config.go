package utils

import (
	"fmt"
	"log"
	"os"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"github.com/koxanybak/todo-api/models"
	"github.com/nats-io/nats.go"
)

// DB is a postgresql database connection
var DB *pg.DB

// NC is a NATS connection
var NC *nats.Conn

// NatsConnect connects to NATS
func NatsConnect() {
	// Nats
	natsAddr := os.Getenv("NATS")
	log.Println("natsAddr:", natsAddr)
	if natsAddr == "" {
		log.Println("No NATS address provided")
		os.Exit(1);
	}
	var err error
	NC, err = nats.Connect(natsAddr)
	if err != nil {
		panic(err)
	}
}

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