package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func saveCount(cnt int) {
	err := ioutil.WriteFile("./files/count.txt", []byte(strconv.Itoa(cnt)), 0644)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Pongs now:", cnt)
	}
}

func main() {
	r := mux.NewRouter()
	cnt := 0
	saveCount(cnt)

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(fmt.Sprintf("pong %d", cnt)))
		cnt++
		saveCount(cnt)
	})

	log.Fatal(http.ListenAndServe("0.0.0.0:8000", r))
}