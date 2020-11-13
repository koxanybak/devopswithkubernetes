package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"time"
)

func main() {
	ticker := time.NewTicker(5 * time.Second)

	curr := time.Now().Local().String()
	err := ioutil.WriteFile("../files/stamp.txt", []byte(time.Now().Local().String()), 0644)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(curr)
	}

	for {
		<- ticker.C
		curr := time.Now().Local().String()
		err := ioutil.WriteFile("../files/stamp.txt", []byte(time.Now().Local().String()), 0644)
		if err != nil {
			log.Println(err)
		} else {
			fmt.Println(curr)
		}
	}
}