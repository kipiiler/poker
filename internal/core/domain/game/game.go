package game

type Game struct {
	BotID   []string
	Started bool
	Ended   bool
	Winner  string
	History []GameState
	Config  GameConfig
	GameID  string
}

func NewGame(botIDs []string, GameID string) *Game {
	return &Game{BotID: botIDs, GameID: GameID, Started: false, Ended: false, Winner: "", History: []GameState{}, Config: *NewDefaultConfig()}

}
func NewGameWithConfig(botIDs []string, GameID string, config *GameConfig) *Game {
	return &Game{BotID: botIDs, GameID: GameID, Started: false, Ended: false, Winner: "", History: []GameState{}, Config: *config}
}
