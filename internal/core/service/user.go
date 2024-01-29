package service

import (
	"errors"
	port "huskyholdem/port"

	"github.com/google/uuid"
)

type UserService struct {
	userRepository port.UserRepository
	userCache      port.UserCache
}

func NewUserService(userRepository port.UserRepository, userCache port.UserCache) *UserService {
	return &UserService{userRepository: userRepository, userCache: userCache}
}

// Login checks if the user exists and if the password is correct.
func (us *UserService) Login(email string, password string) error {
	user, err := us.userRepository.GetUserByUsername(email)
	if err != nil {
		return err
	}
	if user.Password != password {
		return errors.New("Incorrect password")
	}
	return nil
}

func (us *UserService) GenerateAuthToken(email string) {
	// Generate auth token
	authToken := uuid.New().String()

	us.userCache.AddKey(authToken, 10800)
}

func (us *UserService) GenerateBotToken(email string, botId string) {
	// TODO: Check if botId is valid with BotRepository

	// Generate bot token
	botToken := uuid.New().String()

	us.userCache.AddKey(botToken, 10800)
}
