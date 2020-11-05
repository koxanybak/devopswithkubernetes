package main

import (
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/google/uuid"
)

func main() {
	rnd := uuid.New()
	ticker := time.NewTicker(5 * time.Second)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	log.Println(rnd)
	for {
		select {
		case <- ticker.C:
			log.Println(rnd)
		case <- quit:
			ticker.Stop()
			return
		}
	}
}