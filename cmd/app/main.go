package main

import (
	"fmt"
	"log"

	"github.com/runkkaittt/go-tokadb-core/internal/server"
	store "github.com/runkkaittt/go-tokadb-core/internal/store"
)

func main() {
	db := store.NewStore()
	bc1 := db.NewBucket("testBucket")
	if err := bc1.LoadFromFile(); err != nil {
		log.Panic(err)
	}

	bc2 := db.NewBucket("realBucket")
	if err := bc2.LoadFromFile(); err != nil {
		log.Panic(err)
	}

	dbServer := server.New()
	fmt.Println(dbServer.AuthToken)
	dbServer.Start(db)
}
