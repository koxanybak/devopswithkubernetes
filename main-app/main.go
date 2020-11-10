package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func main() {
	rnd := uuid.New()
	ticker := time.NewTicker(5 * time.Second)

	curr := time.Now()
	req := make(chan struct{})
	res := make(chan time.Time)

	log.Println(rnd)
	go func() {
		for {
			select {
			case <- ticker.C:
				log.Println(rnd)
				curr = time.Now()
			case <- req:
				res <- curr
			}
		}
	}()

	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		req <- struct{}{}
		curr := <-res
		w.Write([]byte( fmt.Sprint(curr.Local(), rnd.String()) ))
	})
	log.Fatal(http.ListenAndServe("0.0.0.0:8000", r))
}