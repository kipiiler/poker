package port

import "huskyholdem/user"

type UserCache interface {
	CheckKeyExists(key string) bool
	AddKey(key string, expire int64)
	RemoveKey(key string)
}

type UserRepository interface {
	GetUserByUsername(email string) (*user.User, error)
	GetUserAuthTokens(email string) ([]string, error)
	GetUserBotTokens(email string) ([]string, error)
	GetUserPassword(email string) (string, error)
}

type UserService interface {
	NewUserService()
	Login(email string, password string) error
	GenerateAuthToke(email string)
	GenerateBotToken(email string, botId string)
}
