package api

import (
	"fmt"
	"littleapi/storage"
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
	port  string
	store storage.Storage
}

func NewServer(port string, store storage.Storage) *Server {
	return &Server{
		port:  port,
		store: store,
	}
}

func (s *Server) Start() error {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/user/{id}", s.handleGetUserByID).Methods("GET")
	router.HandleFunc("/user/{id}", s.handleDeleteUserByID).Methods("DELETE")
	router.HandleFunc("/user", s.handleCreateUser).Methods("POST")

	router.Use(contentTypeApplicationJsonMiddleware)
	return http.ListenAndServe(s.port, router)
}

func errorHandler(w http.ResponseWriter, r *http.Request, status int, message string) {
	w.WriteHeader(status)
	if message != "" {
		fmt.Fprint(w, message)
	}
}

func contentTypeApplicationJsonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("X-Corp-Name", "squarehole")

		next.ServeHTTP(w, r)
	})
}
