package port

import "huskyholdem/game"

// GameRepository is the interface that wraps the basic game operations,
// represents the repository layer of the game domain (database interation).
type GameRepository interface {

	// GetGameById returns a game by its id.
	GetGameStateById()
	GetGameById(id string) (*game.Game, error)
	GetAllGame()
	CreateNewGame(botIDs []string, g *game.Game) error
	UpdateGameState(gameID string, gameState *game.GameState) error
	UpdateGame(gameID string, game *game.Game) error
}

// GameService is the interface that wraps the basic game operations.
type GameService interface {
	NewGameService()
	CreateNewGame()
}
