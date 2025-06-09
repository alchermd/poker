package inmemory

import "github.com/alchermd/poker/core"

// SaveSession simply append the given Session to the Game's sessions field.
func (r *Repository) SaveSession(g *core.Game, s *core.Session) (*core.Session, error) {
	g.Sessions = append(g.Sessions, s)
	return s, nil
}
