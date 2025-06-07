package inmemory

import (
	"github.com/alchermd/poker/core"
	"github.com/alchermd/poker/repository"
)

// SaveGame persists the given Game in memory.
func (r *Repository) SaveGame(game *core.Game) (*core.Game, error) {
	if game.ID == 0 {
		// Game is supposed to be created.
		r.currentGameId++
		game.ID = r.currentGameId
		r.games = append(r.games, game)
		return game, nil
	} else {
		// Game is supposed to be updated.
		for idx, g := range r.games {
			if g.ID == game.ID {
				r.games[idx] = game
				return game, nil
			}
		}

		return nil, repository.ErrGameNotFound
	}
}
