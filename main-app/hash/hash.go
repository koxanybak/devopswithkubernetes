package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func main() {
	rnd := uuid.New()

	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		date, err := ioutil.ReadFile("../files/stamp.txt")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		pongs, err := ioutil.ReadFile("../files/count.txt")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		res := fmt.Sprintf("%s: %s\nPing / Pongs: %s", string(date), rnd.String(), string(pongs))
		io.WriteString(w, res)
	})
	log.Fatal(http.ListenAndServe("0.0.0.0:8000", r))
}