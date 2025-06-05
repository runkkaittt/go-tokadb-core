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

	if err := db.LoadFromFile(); err != nil {
		log.Panic(err)
	}
	fmt.Println(db.Data)
}
