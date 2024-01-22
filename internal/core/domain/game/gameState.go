package game

import (
	c "huskyholdem/card"
)

type GameState struct {
	GameStateID string
	GameID      string
	Pot         int
	Bets        map[string]int
	SmallBlind  int
	BigBlind    int
	Middle      []*c.Card
	Folded      map[string]string
	Decision    map[string]string
	Hand        map[string][]*c.Card
	IsRunning   bool
	Turn        int
	Around      int
	Deck        *c.Deck
}

func NewGameState(gameID string, pot int, bets map[string]int, smallBlind int, bigBlind int, middle []*c.Card, folded map[string]string, decision map[string]string, hand map[string][]*c.Card, isRunning bool, turn int, around int, deck *c.Deck) *GameState {
	return &GameState{GameID: gameID, Pot: pot, Bets: bets, SmallBlind: smallBlind, BigBlind: bigBlind, Middle: middle, Folded: folded, Decision: decision, Hand: hand, IsRunning: isRunning, Turn: turn, Around: around, Deck: deck}
}
