package server

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	store "github.com/runkkaittt/go-tokadb-core/internal/store"
)

type DBServer struct {
	router    *mux.Router
	port      string
	dbName    string
	authToken string
}

type field struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func New() *DBServer {
	return &DBServer{
		router:    mux.NewRouter(),
		port:      ":8081",
		dbName:    "database",
		authToken: fmt.Sprint(sha256.Sum256([]byte(fmt.Sprint(time.Now().Unix())))),
	}
}

func (s *DBServer) Start(db *store.Bucket) error {
	s.configureRouter(db)

	log.Println("Server starting...")
	return http.ListenAndServe(s.port, handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "POST"}),
		handlers.AllowedHeaders([]string{"Content-Type"}),
	)(s.router))
}

func (s *DBServer) SetAuthToken(str string) {
	s.authToken = fmt.Sprint(sha256.Sum256([]byte(str)))
}

func (s *DBServer) configureRouter(db *store.Bucket) {
	s.router.HandleFunc(s.dbName+"/set", s.setHandler(db)).Methods("POST")
	s.router.HandleFunc("/get/{key}", s.getHandler(db)).Methods("GET")
	s.router.HandleFunc(s.dbName+"/delete/{key}", s.deleteHandler(db)).Methods("DELETE")
}

func (s *DBServer) setHandler(db *store.Bucket) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("handling set at %s/n", r.URL.Path)

		// au := r.Header.Get("Auth")
		// if au != s.authToken {
		// 	log.Println("request with invalid authentication token")
		// 	return
		// }

		var field field
		if err := json.NewDecoder(r.Body).Decode(&field); err != nil {
			log.Printf("invalid body: %v/n", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		go db.Set(field.Key, field.Value)
		w.WriteHeader(http.StatusOK)
	}
}

func (s *DBServer) getHandler(db *store.Bucket) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("handling get at %s/n", r.URL.Path)

		// au := r.Header.Get("Auth")
		// if au != s.authToken {
		// 	log.Println("request with invalid authentication token")
		// 	return
		// }

		key := mux.Vars(r)["key"]

		val, ok := db.Get(key)
		if !ok {
			log.Printf("invalid URI: %v/n", r.URL.Path)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		log.Println(val)

		field := field{
			Key:   key,
			Value: fmt.Sprintf("%v", val),
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

func (s *DBServer) deleteHandler(db *store.Bucket) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("handling delete at %s/n", r.URL.Path)

		// au := r.Header.Get("Auth")
		// if au != s.authToken {
		// 	log.Println("request with invalid authentication token")
		// 	return
		// }

		key := mux.Vars(r)["key"]

		go db.Delete(key)
		w.WriteHeader(http.StatusOK)
	}
}
