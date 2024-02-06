package adapter

import (
	"database/sql"
	"encoding/json"
	"errors"
	"huskyholdem/game"
	"huskyholdem/gameState"

	pq "github.com/lib/pq"
)

type GameRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *GameRepository {
	return &GameRepository{db: db}
}

func stringToStringMap(Map string) map[string]string {
	m := make(map[string]string)
	json.Unmarshal([]byte(Map), m)
	return m
}

func stringToCardMap(Map string) map[string]Card {
	m := stringToStringMap(Map)
	a := make(map[string]Card)
	for k, v := range m {
		a[k] = Card.MakeCardFromString(v)
	}
	return a
}

func stringToIntMap(Map string) map[string]int {
	m := make(map[string]int)
	json.Unmarshal([]byte(Map), m)
	return m
}

func makeCardArray(arr []string) []Card {
	middle := make([]Card, len(arr))
	for i := range arr {
		middle[i] = arr[i].MakeCardFromString()
	}
	return middle
}

func makeArray(arr []Card) []string {
	middle := make([]string, len(arr))
	for i := range arr {
		middle[i] = arr[i].ToString()
	}
	return middle
}

func MarshalCardMap(Map map[string]Card) (string, error) {
	m := make(map[string]string)
	for k, v := range Map {
		m[k] = v.ToString()
	}
	s, err := json.Marshal(m)

	return string(s), err
}

// GetGameById returns a game by its id.
func (repo *GameRepository) GetGameStateById(id string) (*gameState.GameState, error) {
	sqlCode := `SELECT * FROM gameState WHERE GameStateID=$1;`
	var gameState gameState.GameState
	var Bets string
	var Middle []string
	var Folded string
	var Decision string
	var Hand string
	var Deck []string
	row := repo.db.QueryRow(sqlCode, id)
	err := row.Scan(&gameState.GameStateID, &gameState.GameID, &gameState.Pot, &Bets, &gameState.SmallBlind, &gameState.BigBlind, pq.Array(&Middle), &Folded, &Decision, &Hand, &gameState.IsRunning, &gameState.Turn, &gameState.Around, pq.Array(&Deck))
	gameState.Bets = stringToIntMap(Bets)
	gameState.Middle = makeCardArray(Middle)
	gameState.Folded = stringToStringMap(Folded)
	gameState.Decision = stringToStringMap(Decision)
	gameState.Hand = stringToCardMap(Hand)
	gameState.Deck = makeCardArray(Deck)
	if err == nil {
		return &gameState, nil
	} else if err == sql.ErrNoRows {
		return nil, errors.New("no id found")
	}
	return nil, err
}

func (repo *GameRepository) GetGameById(id string) (*game.Game, error) {
	sqlCode := `SELECT * FROM games WHERE GameID=$1;`
	var game game.Game
	row := repo.db.QueryRow(sqlCode, id)
	err := row.Scan(pq.Array(&game.BotID), &game.Started, &game.Ended, &game.Winner, pq.Array(game.History), &game.Config, &game.GameID)
	if err == nil {
		return &game, nil
	} else if err == sql.ErrNoRows {
		return nil, errors.New("no id found")
	}
	return nil, err
}

func (repo *GameRepository) GetAllGame() (*[]game.Game, error) {
	sqlCode := `SELECT * FROM games;`

	rows, err := repo.db.Query(sqlCode)
	if err != nil {
		return nil, err
	}
	games := []game.Game{}
	defer rows.Close()
	for rows.Next() {
		var game *game.Game
		err2 := rows.Scan(pq.Array(&game.BotID), &game.Started, &game.Ended, &game.Winner, pq.Array(game.History), &game.Config, &game.GameID)
		if err2 != nil {
			return nil, err
		}
		games = append(games, *game)
	}
	return &games, nil

}

func (repo *GameRepository) CreateNewGame(botIDs []string, g *game.Game) error {
	sqlCode := `
	INSERT INTO games (BotID, Started, Ended, Winner, History, Config, GameID)
	VALUES ($1, $2, $3, $4, $5, $6, $7)`
	_, err := repo.db.Exec(sqlCode, pq.Array(g.BotID), g.Started, g.Ended, g.Winner, g.History, g.Config, g.GameID)
	return err
}

func (repo *GameRepository) UpdateGameState(gameID string, gameState *game.GameState) error {
	sqlStatement := `
	UPDATE gameState
	SET GameStateID = $2, GameID = $3, Pot = $4, Bets = $5, SmallBlind = $6, BigBlind = $7, Middle = $8, Folded = $9, Decision = $10, Hand = $11, IsRunning = $12, Turn = $13, Around = $14, Deck = $15
	WHERE id = $1;`
	_, err := repo.db.Exec(sqlStatement, gameID, gameState.GameStateID, gameState.GameID, gameState.Pot, json.Marshal(gameState.Bets), gameState.SmallBlind, gameState.BigBlind, pq.Array(makeArray(gameState.Middle)), json.Marshal(gameState.Folded), json.Marshal(gameState.Decision), MarshalCardMap(gameState.Hand), gameState.IsRunning, gameState.Turn, gameState.Around, pq.Array(makeArray(gameState.Deck.cards)))
	return err
}

func (repo *GameRepository) UpdateGame(gameID string, game *game.Game) error {
	sqlStatement := `
	UPDATE gameState
	SET BotID = $2, Started = $3, Ended = $4, Winner = $5, History = $6, Config = $7, GameID = $8 
	WHERE id = $1;`
	_, err := repo.db.Exec(sqlStatement, gameID, pq.Array(game.BotID), game.Started, gameState.Ended, game.Winner, pq.Array(game.History), game.Config, game.GameID)
	return err
}
