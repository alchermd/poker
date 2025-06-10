package core

import "time"

// Results is a map of a player and their winnings for a session.
type Results map[string]int

// Session represents a single session of a Game.
type Session struct {
	Date    time.Time
	Results Results
}

// NewSession initializes a Session struct.
func NewSession(date time.Time, results Results) *Session {
	s := &Session{
		date,
		results,
	}
	return s
}

// BigWinner calculates which of the session's player has the largest winnings.
func (s *Session) BigWinner() string {
	largestWinnings := 0
	bigWinner := ""
	for player, winning := range s.Results {
		if winning > largestWinnings {
			largestWinnings = winning
			bigWinner = player
		}
	}
	return bigWinner
}
