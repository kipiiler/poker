package bot

type Bot struct {
	ID       string   `json:"id"`
	Name     string   `json:"name"`
	ImgUrl   string   `json:"imgUrl"`
	UserID   string   `json:"userId"`
	BotToken []string `json:"botToken"`
	Keys     []string `json:"keys"`
}

type BotMetaData struct {
	Name   string `json:"name"`
	ImgUrl string `json:"imgUrl"`
}

func NewBotObject(id, name, imgUrl, userID string, botToken []string, keys []string) *Bot {
	return &Bot{
		ID:       id,
		Name:     name,
		ImgUrl:   imgUrl,
		UserID:   userID,
		BotToken: botToken,
		Keys:     keys,
	}
}
