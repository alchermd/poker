package repository

import (
	"errors"
	"github.com/alchermd/poker/core"
)

type Repository interface {
	// SaveGame either creates a new Game (if game.ID == 0) or updates it (if game.ID != 0).
	SaveGame(game *core.Game) (*core.Game, error)
	// GetGame retrieves a Game that matches the given ID.
	GetGame(id int) (*core.Game, error)
}

var ErrGameNotFound = errors.New("game not found")
