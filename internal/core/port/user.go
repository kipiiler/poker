package port

import "huskyholdem/user"

type UserCache interface {
	CheckKeyExists(key string) (bool, error)
	AddKey(key string, expire int64) error
	RemoveKey(key string) error
}

type UserRepository interface {
	GetUserByEmail(email string) (*user.User, error)
	GetUserAuthTokens(email string) ([]string, error)
	GetUserBotTokens(email string) ([]string, error)
	GetUserPassword(email string) (string, error)
	AddUserAuthToken(email string, token string) error
	AddUserBotToken(email string, token string) error
	DeleteUserAuthToken(email string, token string) error
	DeleteUserBotToken(email string, token string) error
}

type UserService interface {
	NewUserService()
	Login(email string, password string) error
	GenerateAuthToken(email string) (string, error)
	GenerateBotToken(email string, botId string)
	CheckAuthToken(email string, token string) (bool, error)
}
