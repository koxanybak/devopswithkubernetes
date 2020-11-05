package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/koxanybak/devopswithkubernetes/todo/routes"
)

func main() {
	godotenv.Load()
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("Error: no port specified")
	}
	r := mux.NewRouter()

	spa := routes.SpaHandler{StaticPath: "build", IndexPath: "index.html"}
	r.PathPrefix("/").Handler(spa)

	srv := &http.Server{
		Addr: fmt.Sprintf("0.0.0.0:%s", port),
		Handler: r,
	}

	fmt.Println("Server started in port", port)
	log.Fatal(srv.ListenAndServe())
}