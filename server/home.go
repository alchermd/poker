package server

import (
	"net/http"
)

const tplHome = "home.gohtml"

// HomeHandler serves the home page content.
func (s *Server) HomeHandler(w http.ResponseWriter, r *http.Request) {
	err := s.templates[tplHome].ExecuteTemplate(w, tplHome, nil)
	if err != nil {
		s.l.Printf("Error writing response %q", err)
		http.Error(w, "Error writing response", http.StatusInternalServerError)
		return
	}
}
