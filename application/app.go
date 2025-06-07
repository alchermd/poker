package application

import (
	"github.com/alchermd/poker/repository"
)

// App encapsulates the application-level services.
type App struct {
	repo repository.Repository
}

// NewApp creates a new App
func NewApp(repo repository.Repository) *App {
	app := &App{
		repo,
	}
	return app
}
