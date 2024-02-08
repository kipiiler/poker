package adapters

import (
	"database/sql"
	"errors"
	"huskyholdem/bot"

	"github.com/lib/pq"
)

type BotRepository struct {
	db *sql.DB
}

func NewBotRepository(db *sql.DB) *BotRepository {
	return &BotRepository{db: db}
}

func (repo *BotRepository) GetBotByID(id string) (*bot.Bot, error) {
	sqlCode := `SELECT * FROM bots WHERE bot_id=$1;`
	var bot bot.Bot
	row := repo.db.QueryRow(sqlCode, id)
	err := row.Scan(&bot.ID, &bot.Name, &bot.ImgUrl, &bot.UserID, pq.Array(&bot.BotToken), pq.Array(&bot.Keys))
	if err == nil {
		return &bot, nil
	} else if err == sql.ErrNoRows {
		return nil, errors.New("no bot found")
	}
	return nil, err
}

func (repo *BotRepository) GetBotByUserID(userID string) ([]*bot.Bot, error) {
	sqlCode := `SELECT * FROM bots WHERE user_id=$1;`
	rows, err := repo.db.Query(sqlCode, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var bots []*bot.Bot
	for rows.Next() {
		var bot bot.Bot
		err = rows.Scan(&bot.ID, &bot.Name, &bot.ImgUrl, &bot.UserID, pq.Array(&bot.BotToken), pq.Array(&bot.Keys))
		if err != nil {
			return nil, err
		}
		bots = append(bots, &bot)
	}
	return bots, nil
}

func (repo *BotRepository) CreateNewBot(bot *bot.Bot) error {
	sqlCode := `INSERT INTO bots (bot_id, bot_name, img_url, user_id, bot_tokens, keys) VALUES ($1, $2, $3, $4, $5, $6);`
	_, err := repo.db.Exec(sqlCode, bot.ID, bot.Name, bot.ImgUrl, bot.UserID, pq.Array(bot.BotToken), pq.Array(bot.Keys))
	return err
}

func (repo *BotRepository) UpdateBot(bot *bot.Bot) error {
	sqlCode := `UPDATE bots SET bot_name=$1, img_url=$2, bot_tokens=$3, keys=$4 WHERE bot_id=$5;`
	_, err := repo.db.Exec(sqlCode, bot.Name, bot.ImgUrl, pq.Array(bot.BotToken), pq.Array(bot.Keys), bot.ID)
	return err
}

func (repo *BotRepository) AddKey(botId string, key string) error {
	sqlCode := `
	UPDATE bots
	SET keys = array_append(keys, $1)
	WHERE bot_id = $2;`

	_, err := repo.db.Exec(sqlCode, key, botId)
	return err
}

func (repo *BotRepository) RemoveKey(botId string, key string) error {
	sqlCode := `
	UPDATE bots
	SET keys = array_remove(keys, $1)
	WHERE bot_id = $2;`

	_, err := repo.db.Exec(sqlCode, key, botId)
	return err
}

func (repo *BotRepository) AddBotToken(botId string, token string) error {
	sqlCode := `
	UPDATE bots
	SET bot_tokens = array_append(bot_tokens, $1)
	WHERE bot_id = $2;`

	_, err := repo.db.Exec(sqlCode, token, botId)
	return err
}

func (repo *BotRepository) RemoveBotToken(botId string, token string) error {
	sqlCode := `
	UPDATE bots
	SET bot_tokens = array_remove(bot_tokens, $1)
	WHERE bot_id = $2;`

	_, err := repo.db.Exec(sqlCode, token, botId)
	return err
}
