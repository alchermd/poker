package server

import (
	"net/http"
)

const tplHome = "home.gohtml"

// Home serves the home page content.
func (s *Server) Home(w http.ResponseWriter, r *http.Request) {
	err := s.templates.ExecuteTemplate(w, tplHome, nil)
	if err != nil {
		s.l.Printf("Error writing response %q", err)
	}
}
