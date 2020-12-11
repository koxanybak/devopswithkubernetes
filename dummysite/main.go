package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func serveHTML(htmlData []byte) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "text/html; charset=UTF-8")
		w.Write(htmlData)
	}
}

func main() {
	dummysite := os.Getenv("WEBSITE_URL")
	if (dummysite == "") {
		panic("No website url provided")
	}

	resp, err := http.Get(dummysite)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	htmlData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	r := mux.NewRouter()
	r.HandleFunc("/", serveHTML(htmlData))

	log.Fatal(http.ListenAndServe("0.0.0.0:8000", r))
}