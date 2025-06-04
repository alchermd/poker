package server

import (
	"net/http"
)

// RegisterRoutes attaches handlers to the given router.
// We can treat this method as the central hub of all the routes available in the system.
// Individual handler functions are found in their own files.
func (s *Server) RegisterRoutes(router *http.ServeMux) *http.ServeMux {
	router.HandleFunc("/", s.Home)
	return router
}
