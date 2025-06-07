package main

import (
	"github.com/alchermd/poker/application"
	"github.com/alchermd/poker/repository/inmemory"
	"github.com/alchermd/poker/server"
)

func main() {
	repo := inmemory.NewRepo()
	app := application.NewApp(repo)
	s := server.NewServer(app)
	s.Run()
}
