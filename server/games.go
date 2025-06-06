package server

import "net/http"

const tplCreateGame = "create-game.gohtml"

func (s *Server) CreateGame(w http.ResponseWriter, r *http.Request) {
	err := s.templates[tplCreateGame].ExecuteTemplate(w, tplCreateGame, nil)
	if err != nil {
		s.l.Printf("Error writing response %q", err)
	}
}
