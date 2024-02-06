package port

import "huskyholdem/bot"

type BotCache interface {
	CheckKeyExists(key string) (bool, error)
	AddKey(key string, value string) error
	AddKeyWithExpiration(key string, value string, expiration int) error
	GetKey(key string) (string, error)
	RemoveKey(key string) error
	GetKeysWithPrefix(prefix string) ([]string, error)
}

type BotRepository interface {
	GetBotByID(id string) (*bot.Bot, error)
	GetBotByUserID(userID string) ([]*bot.Bot, error)
	CreateNewBot(bot *bot.Bot) error
	UpdateBot(bot *bot.Bot) error
	AddKey(botId string, key string) error
	RemoveKey(botId string, key string) error
	AddBotToken(botId string, token string) error
	RemoveBotToken(botId string, token string) error
}

type BotService interface {
	NewBotService()
	GetBotByID(id string) (*bot.Bot, error)
	GetBotByUserID(userID string) ([]*bot.Bot, error)
	CreateNewBot(name string, imgUrl string, email string) (string, error)
	UpdateBotMetadata(id string, botData *bot.BotMetaData) error
	GetKeyValuesFromCache(id string) (map[string]string, error)
	GetKeyFromCache(id string, key string) (string, error)
	AddKeyValuesToCache(id string, key string, value string) error
	RemoveKeyValueFromCache(id string, key string) error
	FlushCache(id string) error
	CheckBotToken(botId string, email string) (bool, error)
	GenerateBotToken(botId string) (string, error)
	// IsUserOwnBot(email string, botId string) (bool, error)
}
