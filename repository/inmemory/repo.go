package inmemory

import (
	"github.com/alchermd/poker/core"
)

// Repository is an in-memory data store.
type Repository struct {
	currentGameId int
	games         []*core.Game
}

// NewRepo initializes a new Repository.
func NewRepo() *Repository {
	r := &Repository{
		0,
		[]*core.Game{},
	}
	return r
}
