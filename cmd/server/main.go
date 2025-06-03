package main

import "github.com/alchermd/poker/server"

func main() {
	s := server.NewServer()
	s.Run()
}
