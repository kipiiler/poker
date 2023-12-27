package port

// GameRepository is the interface that wraps the basic game operations,
// represents the repository layer of the game domain (database interation).
type GameRepository interface {

	// GetGameById returns a game by its id.
	GetGameStateById()
	GetGameId()
	GetAllGame()
	CreateNewGame()
	StartGame()
	EndGame()
	PauseGame()
	ResumeGame()
	UpdateGameState()
}

// GameService is the interface that wraps the basic game operations.
type GameService interface {
	NewGameService()
	CreateNewGame()
}
