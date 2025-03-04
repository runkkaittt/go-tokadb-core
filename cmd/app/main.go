package main

import (
	"fmt"
	"log"

	store "github.com/runkkaittt/go-tokadb-core/internal/store"
)

func main() {
	db := store.NewBucket("testBucket")
	if err := db.LoadFromFile(); err != nil {
		log.Panic(err)
	}

	db.Set("test key 5", true)
	db.Set("test key 4", "hello toka db")

	val1, _ := db.Get("test key 1")
	fmt.Println(val1)

	_, ok := db.Get("aksjdh")
	fmt.Println(ok)

	if err := db.SaveToFile(); err != nil {
		log.Panic(err)
	}
}
