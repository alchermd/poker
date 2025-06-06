package server

import (
	"github.com/gorilla/mux"
)

// RegisterRoutes attaches handlers to the given router.
// We can treat this method as the central hub of all the routes available in the system.
// Individual handler functions are found in their own files.
func (s *Server) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/", s.Home).Methods("GET")
	router.HandleFunc("/create", s.CreateGame).Methods("GET", "POST")
}
