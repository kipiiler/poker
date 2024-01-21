package ports

type GameActions interface {
    HandleFold(clientID string) error
    HandleCheck(clientID string) error
    HandleCall(clientID string) error
    HandleRaise(clientID string, amount int) error
}

type GameStateStore interface {
    GetGameState(gameID string) (GameState, error)
    SaveGameState(gameID string, gameState GameState) error
}

type Cache interface {
    Store(key string, value interface{}) error
    Get(key string) (interface{}, error)
}





