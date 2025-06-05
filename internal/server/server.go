package server

import (
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type DBServer struct {
	router *mux.Router
	port   string
	dbName string
}

func New() *DBServer {
	return &DBServer{
		router: mux.NewRouter(),
		port:   ":8081",
		dbName: "database",
	}
}

func (s *DBServer) Start() error {
	s.configureRouter()

	return http.ListenAndServe(s.port, handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "POST"}),
		handlers.AllowedHeaders([]string{"Content-Type"}),
	)(s.router))
}

func (s *DBServer) configureRouter() {
	s.router.HandleFunc(s.dbName+"/set", s.setHandler).Methods("POST")
	s.router.HandleFunc(s.dbName+"/get/{key}", s.getHandler).Methods("GET")
	s.router.HandleFunc(s.dbName+"/delete/{key}", s.deleteHandler).Methods("DELETE")
}

func (s *DBServer) setHandler(w http.ResponseWriter, req *http.Request) {}

func (s *DBServer) getHandler(w http.ResponseWriter, req *http.Request) {}

func (s *DBServer) deleteHandler(w http.ResponseWriter, req *http.Request) {}
