package server

import (
	"errors"
	"github.com/alchermd/poker/core"
	"github.com/alchermd/poker/repository"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type createSessionPayload struct {
	shouldIncludePlayer []string
	player              []string
	result              []int
}

// CreateSessionHandler handles requests for creation of a new Session.
// The expectation is that this will be called asynchronously from the frontend.
func (s *Server) CreateSessionHandler(w http.ResponseWriter, r *http.Request) {
	// Check if the Game ID is of correct format.
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		s.l.Printf("Cannot convert ID to int %q", err)
		http.Error(w, "Game not found", http.StatusNotFound)
		return
	}

	// Check if the Game ID is existent.
	g, err := s.app.GetGame(id)
	if errors.Is(err, repository.ErrGameNotFound) {
		s.l.Printf("Game not found %q", err)
		http.Error(w, "Game not found", http.StatusNotFound)
		return
	}

	// Parse form data
	// Make sure all results can be converted to int.
	err = r.ParseForm()
	var resultsInt []int
	for _, rawResult := range r.PostForm["result"] {
		result, err := strconv.Atoi(rawResult)
		if err != nil {
			s.l.Printf("Cannot convert result to int %q", err)
			http.Error(w, "Invalid result", http.StatusBadRequest)
			return
		}
		resultsInt = append(resultsInt, result)
	}

	// Prepare payload struct.
	payload := &createSessionPayload{
		shouldIncludePlayer: r.PostForm["shouldIncludePlayer"],
		player:              r.PostForm["player"],
		result:              resultsInt,
	}

	// Convert payload to Results
	var results core.Results
	for idx, shouldIncludePlayer := range payload.shouldIncludePlayer {
		if shouldIncludePlayer != "checked" {
			continue
		}
		player := payload.player[idx]
		result := payload.result[idx]
		results[player] = result
	}

	// Call a service function.
	_, err = s.app.CreateSession(g, results)
	if err != nil {
		s.l.Printf("Creating session failed %q", err)
		http.Error(w, "Creating session failed", http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
