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
	mainRouter := mux.NewRouter()

	spa := routes.SpaHandler{StaticPath: "build", IndexPath: "index.html"}
	routes.APIRouter(mainRouter)

	mainRouter.PathPrefix("/").Handler(spa)

	srv := &http.Server{
		Addr: fmt.Sprintf("0.0.0.0:%s", port),
		Handler: mainRouter,
	}

	fmt.Println("Server started in port", port)
	log.Fatal(srv.ListenAndServe())
}