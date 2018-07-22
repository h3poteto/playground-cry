package main

import (
	"log"
	"time"
)

func main() {
	for {
		log.Println("cry")
		time.Sleep(10 * time.Second)
	}
}
