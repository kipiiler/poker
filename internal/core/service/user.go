package service

import (
	"errors"
	"fmt"
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
	user, err := us.userRepository.GetUserByEmail(email)
	if err != nil {
		return err
	}
	if user.Password != password {
		return errors.New("Incorrect password")
	}
	return nil
}

func (us *UserService) GenerateAuthToken(email string) (string, error) {
	// Generate auth token
	userClaims := NewUserClaims(email)
	authToken, errToken := NewAuthToken(userClaims)
	if errToken != nil {
		return "", errToken
	}

	err := us.userCache.AddKey(authToken, 10800)
	if err != nil {
		return "", err
	}

	fmt.Println("authToken: ", authToken)

	err2 := us.userRepository.AddUserAuthToken(email, authToken)
	if err2 != nil {
		return "", err2
	}

	return authToken, nil
}

func (us *UserService) GenerateBotToken(email string, botId string) {
	// TODO: Check if botId is valid with BotRepository

	// Generate bot token
	botToken := uuid.New().String()

	err := us.userCache.AddKey(botToken, 10800)
	if err != nil {
		us.userRepository.AddUserBotToken(email, botToken)
	}
}
