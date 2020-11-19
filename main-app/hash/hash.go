package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

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
		
		pongCountRes, err := http.Get("http://pingpong-svc/count")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer pongCountRes.Body.Close()
		buf := new(strings.Builder)
		_, err = io.Copy(buf, pongCountRes.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		pongs := buf.String()

		res := fmt.Sprintf(
			"%s\n%s: %s\nPing / Pongs: %s",
			os.Getenv("MESSAGE"),
			string(date),
			rnd.String(),
			string(pongs),
		)
		io.WriteString(w, res)
	})
	log.Fatal(http.ListenAndServe("0.0.0.0:8000", r))
}