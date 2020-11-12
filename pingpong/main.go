package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	cnt := 0

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(fmt.Sprintf("pong %d", cnt)))
		cnt++
	})

	log.Fatal(http.ListenAndServe("0.0.0.0:8000", r))
}