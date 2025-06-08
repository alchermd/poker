package server

import (
	"errors"
	"fmt"
	"github.com/alchermd/poker/repository"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

const tplCreateGame = "create-game.gohtml"

// Think -- should this be defined in the server layer or in the core layer?
type createGamePayload struct {
	name        string
	description string
	players     []string
}

// CreateGameHandler serves the HTML form for creating a game and the POST handling of the submitted data.
func (s *Server) CreateGameHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		err := s.templates[tplCreateGame].ExecuteTemplate(w, tplCreateGame, nil)
		if err != nil {
			s.l.Printf("Error writing response %q", err)
			http.Error(w, "Error writing response", http.StatusInternalServerError)
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
		game, err := s.app.CreateGame(payload.name, payload.description, payload.players)
		if err != nil {
			s.l.Printf("Error creating game %q", err)
			http.Error(w, "Error creating game", http.StatusBadRequest)
			return
		}

		http.Redirect(w, r, fmt.Sprintf("/games/%d", game.ID), http.StatusFound)
		return
	}
}

const tplShowGame = "show-game.gohtml"

// ShowGameHandler serves the details page of a specific Game.
func (s *Server) ShowGameHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		s.l.Printf("Cannot convert ID to int %q", err)
		http.Error(w, "Game not found", http.StatusNotFound)
		return
	}

	g, err := s.app.GetGame(id)
	if errors.Is(err, repository.ErrGameNotFound) {
		s.l.Printf("Game not found %q", err)
		http.Error(w, "Game not found", http.StatusNotFound)
		return
	}

	err = s.templates[tplShowGame].ExecuteTemplate(w, tplShowGame, g)
	if err != nil {
		s.l.Printf("Error writing response %q", err)
		http.Error(w, "Error writing response", http.StatusInternalServerError)
		return
	}
}
