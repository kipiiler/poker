package domain

type Game struct {
    botID []string
    started string
    ended string
    winner string
    history []GameState
    config GameConfig
}

