package routes

import (
	"github.com/gorilla/mux"
	"github.com/koxanybak/devopswithkubernetes/todo/routes/api"
)

// APIRouter returns a router for the REST-api
func APIRouter(mainRouter *mux.Router) {
	r := mainRouter.PathPrefix("/image").Subrouter()
	r.HandleFunc("/", api.ImageHandler)
}