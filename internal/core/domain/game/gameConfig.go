package game

type GameConfig struct {
	TimePerRound int `json:"timePerRound"`
}

func NewGameConfig(timePerRound int) *GameConfig {
	return &GameConfig{TimePerRound: timePerRound}
}

func NewDefaultConfig() *GameConfig {
	return &GameConfig{TimePerRound: 30}
}
