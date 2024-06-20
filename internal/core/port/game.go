package port

import "huskyholdem/game"

// GameRepository is the interface that wraps the basic game operations,
// represents the repository layer of the game domain (database interation).
type GameRepository interface {

	// GetGameById returns a game by its id.
	// GetGameStateById()
	GetGameByID(id string) (*game.Game, error)
	// GetAllGame()
	CreateNewGame(g *game.Game) error
	// UpdateGameState(gameID string, gameState *game.GameState) error
	// UpdateGame(gameID string, game *game.Game) error
}

// GameService is the interface that wraps the basic game operations.
type GameService interface {
	NewGameService(repo GameRepository) *GameService
	CreateNewGame(botIDs []string) (*game.Game, error)
	CreateNewGameWithID(botIDs []string, gameID string) (*game.Game, error)
	GetGameByID(id string) (*game.Game, error)
}
