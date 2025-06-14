package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	store "github.com/runkkaittt/go-tokadb-core/internal/store"
)

func (s *DBServer) setHandler(db *store.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("handling set at %s/n", r.URL.Path)

		// au := r.Header.Get("Auth")
		// if au != s.AuthToken {
		// 	log.Println("request with invalid authentication token")
		// 	return
		// }

		bc := mux.Vars(r)["bucket"]

		var field field
		if err := json.NewDecoder(r.Body).Decode(&field); err != nil {
			log.Printf("invalid body: %v/n", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		go db.Buckets[bc].Set(field.Key, field.Value)
		w.WriteHeader(http.StatusOK)
	}
}

func (s *DBServer) getHandler(db *store.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("handling get at %s/n", r.URL.Path)

		// au := r.Header.Get("Auth")
		// if au != s.AuthToken {
		// 	log.Println("request with invalid authentication token")
		// 	return
		// }

		bc := mux.Vars(r)["bucket"]
		key := mux.Vars(r)["key"]

		val, ok := db.Buckets[bc].Get(key)
		if !ok {
			log.Printf("invalid URI: %v/n", r.URL.Path)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		log.Println(val)

		field := field{
			Key:   key,
			Value: val,
		}

		w.Header().Set("Content-Type", "application/json")

		if err := json.NewEncoder(w).Encode(&field); err != nil {
			log.Printf("invalid URI: %v/n", r.URL.Path)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

func (s *DBServer) deleteHandler(db *store.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("handling delete at %s/n", r.URL.Path)

		// au := r.Header.Get("Auth")
		// if au != s.AuthToken {
		// 	log.Println("request with invalid authentication token")
		// 	return
		// }

		bc := mux.Vars(r)["bucket"]
		key := mux.Vars(r)["key"]

		go db.Buckets[bc].Delete(key)
		w.WriteHeader(http.StatusOK)
	}
}
