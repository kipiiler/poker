package domain

import (
	c "huskyholdem/card"
)

type GameState struct {
	GameID     string
	Pot        int
	Bets       map[string]int
	SmallBlind int
	BigBlind   int
	Middle     []*c.Card
	Folded     map[string]string
	Decision   map[string]string
	Hand       map[string][]*c.Card
	IsRunning  bool
	Turn       int
	Around     int
	Deck       *c.Deck
}
