package server

import (
	"net/http"
)

const tplCreateGame = "create-game.gohtml"

// Think -- should this be defined in the server layer or in the core layer?
type createGamePayload struct {
	name        string
	description string
	players     []string
}

// CreateGame serves the HTML form for creating a game and the POST handling of the submitted data.
func (s *Server) CreateGame(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		err := s.templates[tplCreateGame].ExecuteTemplate(w, tplCreateGame, nil)
		if err != nil {
			s.l.Printf("Error writing response %q", err)
		}
		return
	} else if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			s.l.Printf("Error parsing form %q", err)
			http.Error(w, "Error parsing form", http.StatusBadRequest)
			return
		}

		// TODO: Pass this to a service function to facilitate game creation.
		payload := &createGamePayload{
			name:        r.Form.Get("name"),
			description: r.Form.Get("description"),
			players:     r.Form["players"],
		}
		s.l.Printf("Parsed payload: %#v", payload)

		http.Redirect(w, r, "/games/1", http.StatusFound)
		return
	}
}
