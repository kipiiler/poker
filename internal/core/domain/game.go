package domain

type Game struct {
	BotID   []string
	Started string
	Ended   string
	Winner  string
	History []GameState
	Config  GameConfig
}
