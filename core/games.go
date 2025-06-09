package core

// Game represents a game of Poker.
type Game struct {
	ID          int
	Name        string
	Description string
	Players     []string
	Sessions    []*Session
}

// NewGame initializes a new Game.
func NewGame(id int, name string, description string, players []string) *Game {
	g := &Game{
		id,
		name,
		description,
		players,
		[]*Session{},
	}
	return g
}
