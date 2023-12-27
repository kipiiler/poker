package service

import (
	"huskyholdem/port"
)

// GameService implements port.GameService interface
// and provides an access to the game repository
type GameService struct {
	repo port.GameRepository
}
