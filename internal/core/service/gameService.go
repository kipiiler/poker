package service

import (
	"huskyholdem/card"
	game "huskyholdem/game"
	port "huskyholdem/port"

	"github.com/google/uuid"
)

// GameService implements port.GameService interface
// and provides an access to the game repository
type GameService struct {
	repo    port.GameRepository
	deckMap map[string]*card.Deck
}

// NewGameService creates a new instance of GameService
func NewGameService(repo port.GameRepository) *GameService {
	return &GameService{
		repo:    repo,
		deckMap: map[string]*card.Deck{},
	}
}

// GetGameByID returns a game by its ID
func (s *GameService) CreateNewGame(botIDs []string) (*game.Game, error) {
	gameID := uuid.New().String()
	game := game.NewGame(botIDs, gameID)
	err := s.repo.CreateNewGame(botIDs, game)
	if err != nil {
		return nil, err
	}
	return game, nil
}

func (s *GameService) CreateNewGameWithConfig(botIDs []string, gameConfig *game.GameConfig) (*game.Game, error) {
	if gameConfig == nil {
		gameConfig = game.NewDefaultConfig()
	}
	gameID := uuid.New().String()
	game := game.NewGameWithConfig(botIDs, gameID, gameConfig)
	err := s.repo.CreateNewGame(botIDs, game)
	if err != nil {
		return nil, err
	}
	return game, nil
}

func (s *GameService) FindGameByID(id string) (*game.Game, error) {
	game, err := s.repo.GetGameById(id)
	if err != nil {
		return nil, err
	}
	return game, nil
}

func (s *GameService) StartGame(id string) (*game.Game, error) {
	game, err := s.repo.GetGameById(id)
	if err != nil {
		return nil, err
	}
	game.Started = true
	err = s.repo.UpdateGame(id, game)
	if err != nil {
		return nil, err
	}
	s.deckMap[id] = card.NewDeck()
	return game, nil
}
