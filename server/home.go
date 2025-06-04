package server

import (
	"fmt"
	"net/http"
)

// Home serves the home page content.
func (s *Server) Home(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprint(w, "hello, world!")
	if err != nil {
		s.l.Printf("Error writing response %q", err)
	}
}
