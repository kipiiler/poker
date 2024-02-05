package service

import (
	"huskyholdem/bot"
	port "huskyholdem/port"

	"github.com/google/uuid"
)

type BotService struct {
	botRepository port.BotRepository
	botCache      port.BotCache
}

func NewBotService(botRepository port.BotRepository, botCache port.BotCache) *BotService {
	return &BotService{botRepository: botRepository, botCache: botCache}
}

func (bs *BotService) GetBotByID(id string) (*bot.Bot, error) {
	return bs.botRepository.GetBotByID(id)
}

func (bs *BotService) UpdateBotMetadata(id string, botData *bot.BotMetaData) error {
	bot, err := bs.botRepository.GetBotByID(id)
	if err != nil {
		return err
	}

	bot.Name = botData.Name
	bot.ImgUrl = botData.ImgUrl

	return bs.botRepository.UpdateBot(bot)
}

func (bs *BotService) GetBotByUserID(userID string) ([]*bot.Bot, error) {
	return bs.botRepository.GetBotByUserID(userID)
}

func (bs *BotService) CreateNewBot(name string, imgUrl string, email string) (string, error) {
	id := uuid.New().String()
	bot := bot.NewBotObject(id, name, imgUrl, email, []string{}, []string{})

	return id, bs.botRepository.CreateNewBot(bot)
}

func (bs *BotService) AddKeyValuesToCache(id string, key string, value string) error {
	keyToRedis := id + ":" + key

	err := bs.botCache.AddKey(keyToRedis, value)
	if err != nil {
		return err
	}

	errRepo := bs.botRepository.AddKey(id, key)
	if errRepo != nil {
		return errRepo
	}

	return nil
}

func (bs *BotService) GetKeyFromCache(id string, key string) (string, error) {
	// return bs.botCache.GetKey(id, key)
	keyToRedis := id + ":" + key
	_, err := bs.botCache.CheckKeyExists(keyToRedis)
	if err != nil {
		errRepo := bs.botRepository.RemoveKey(id, key)
		if errRepo != nil {
			return "", errRepo
		}
		return "", err
	}

	return bs.botCache.GetKey(keyToRedis)
}

func (bs *BotService) RemoveKeyValueFromCache(id string, key string) error {
	keyToRedis := id + ":" + key
	err := bs.botCache.RemoveKey(keyToRedis)
	if err != nil {
		return err
	}

	errRepo := bs.botRepository.RemoveKey(id, key)
	if errRepo != nil {
		return errRepo
	}

	return nil
}

func (bs *BotService) GetKeyValuesFromCache(id string) (map[string]string, error) {
	prefix := id + ":"
	keys, err := bs.botCache.GetKeysWithPrefix(prefix)
	if err != nil {
		return nil, err
	}

	bot, err := bs.botRepository.GetBotByID(id)
	if err != nil {
		return nil, err
	}

	if len(bot.Keys) != len(keys) {
		bot.Keys = keys
		err = bs.botRepository.UpdateBot(bot)
		if err != nil {
			return nil, err
		}
	}

	keyValues := make(map[string]string)
	for _, key := range keys {
		value, err := bs.botCache.GetKey(key)
		if err != nil {
			return nil, err
		}
		keyValues[key] = value
	}

	return keyValues, nil
}

func (bs *BotService) FlushCache(id string) error {
	prefix := id + ":"
	keys, err := bs.botCache.GetKeysWithPrefix(prefix)

	if err != nil {
		return err
	}

	for _, key := range keys {
		err := bs.botCache.RemoveKey(key)
		if err != nil {
			return err
		}
	}

	bot, err := bs.botRepository.GetBotByID(id)
	if err == nil {
		return err
	}

	bot.Keys = []string{}
	err = bs.botRepository.UpdateBot(bot)
	if err != nil {
		return err
	}

	return nil
}
