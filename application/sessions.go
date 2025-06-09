package application

import (
	"github.com/alchermd/poker/core"
	"time"
)

// CreateSession creates a new Session and attaches it to the Game.
func (a *App) CreateSession(g *core.Game, results core.Results) (*core.Session, error) {
	s := core.NewSession(time.Now(), results)
	s, err := a.repo.SaveSession(g, s)
	return s, err
}
