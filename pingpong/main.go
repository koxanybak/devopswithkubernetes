package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"github.com/gorilla/mux"
)

// Count is count
type Count struct {
	ID		int64
	Cnt		int64		`pg:",notnull,use_zero"`
}

var db *pg.DB

func saveCount(cnt int64) {
	_, err := db.Model(&Count{ID: 1, Cnt: cnt}).
		OnConflict("(id) DO UPDATE").
		Set("cnt = EXCLUDED.cnt").
		Insert()
	if err != nil {
        fmt.Println("Error inserting:", err)
    }
}

func getCount() int64 {
	var cnts []Count
	var cnt int64
	err := db.Model(&cnts).Where("id = 1").Select()
	if err != nil {
        fmt.Println("Error getting:", err)
		return 0
    }
	if len(cnts) == 0 {
		cnt = 0
		saveCount(0)
	} else {
		cnt = cnts[0].Cnt
	}

	return cnt
}

func main() {
	log.Println(os.Getenv("POSTGRES_PASSWORD"))
	log.Println(os.Getenv("POSTGRES_DB"))
	log.Println(os.Getenv("POSTGRES_USER"))
	db = pg.Connect(&pg.Options{
		User: os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		Database: os.Getenv("POSTGRES_DB"),
		Addr: "postgres-svc:5432",
	})

	defer db.Close()
	for _, model := range []interface{}{&Count{}} {
		err := db.Model(model).CreateTable(&orm.CreateTableOptions{
			IfNotExists: true,
		})
		if err != nil {
			panic(err)
		}
	}


	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		cnt := getCount()
		w.Write([]byte(fmt.Sprintf("pong %d", cnt)))
		saveCount(cnt+1)
	})
	r.HandleFunc("/count", func(w http.ResponseWriter, r *http.Request) {
		cnt := getCount()
		io.WriteString(w, strconv.FormatInt(cnt, 10))
	})

	log.Println("Listeting on port", "8000")
	log.Fatal(http.ListenAndServe(fmt.Sprintf("0.0.0.0:%s", "8000"), r))
}