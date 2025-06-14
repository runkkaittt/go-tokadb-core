package server

import (
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
	AuthToken string
}

type field struct {
	Key   string `json:"key"`
	Value any    `json:"value"`
}

func New() *DBServer {
	return &DBServer{
		router:    mux.NewRouter(),
		port:      ":8081",
		AuthToken: createAuthToken(fmt.Sprint(time.Now().Unix())),
	}
}

func (s *DBServer) Start(db *store.Store) error {
	s.configureRouter(db)

	log.Println("Server starting...")
	return http.ListenAndServe(s.port, handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "POST", "DELETE"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Auth"}),
	)(s.router))
}

func (s *DBServer) configureRouter(db *store.Store) {
	s.router.HandleFunc("/api/{bucket}/set", s.setHandler(db)).Methods("POST")
	s.router.HandleFunc("/api/{bucket}/get/{key}", s.getHandler(db)).Methods("GET")
	s.router.HandleFunc("/api/{bucket}/delete/{key}", s.deleteHandler(db)).Methods("DELETE")
}
