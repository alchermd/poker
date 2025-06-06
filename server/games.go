package server

import "net/http"

const tplCreateGame = "create-game.gohtml"

func (s *Server) CreateGame(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		err := s.templates[tplCreateGame].ExecuteTemplate(w, tplCreateGame, nil)
		if err != nil {
			s.l.Printf("Error writing response %q", err)
		}
		return
	} else if r.Method == http.MethodPost {
		http.Error(w, "TODO: Implement me!", http.StatusNotImplemented)
		return
	}
}
