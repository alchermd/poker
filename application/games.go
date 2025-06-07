package application

import "github.com/alchermd/poker/core"

// CreateGame takes the parameters, constructs a Game instance,
// and passes it along to the configured repository for persistence.
func (a *App) CreateGame(name string, description string, players []string) (*core.Game, error) {
	g := core.NewGame(0, name, description, players)
	g, err := a.repo.SaveGame(g)
	return g, err
}
