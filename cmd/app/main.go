package main

import (
	"log"

	"github.com/runkkaittt/go-tokadb-core/internal/server"
	store "github.com/runkkaittt/go-tokadb-core/internal/store"
)

func main() {
	db := store.NewBucket("testBucket")
	if err := db.LoadFromFile(); err != nil {
		log.Panic(err)
	}

	dbServer := server.New()
	dbServer.Start(db)
}
