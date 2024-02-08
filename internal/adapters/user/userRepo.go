package adapters

import (
	"database/sql"
	"errors"
	"huskyholdem/user"

	pq "github.com/lib/pq"
)

type UserRepository struct {
	db *sql.DB
}

type User struct {
	Email      string
	Password   string
	AuthTokens []string
	BotTokens  []string
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (repo *UserRepository) GetUserByEmail(email string) (*user.User, error) {
	sqlCode := `SELECT * FROM users WHERE email=$1;`
	var user user.User
	row := repo.db.QueryRow(sqlCode, email)
	err := row.Scan(&user.Email, &user.Password, pq.Array(&user.AuthTokens), pq.Array(&user.BotTokens))
	if err == nil {
		return &user, nil
	} else if err == sql.ErrNoRows {
		return nil, errors.New("no email found")
	}
	return nil, err
}

func (repo *UserRepository) GetUserAuthTokens(email string) ([]string, error) {
	sqlCode := `SELECT auth_tokens FROM users WHERE email=$1;`
	var authTokens pq.StringArray = pq.StringArray{}
	row := repo.db.QueryRow(sqlCode, email)
	err := row.Scan(&authTokens)
	AuthTokens := []string(authTokens)
	if err == nil {
		return AuthTokens, nil
	} else if err == sql.ErrNoRows {
		return nil, errors.New("no email found")
	}
	return nil, err
}
func (repo *UserRepository) GetUserBotTokens(email string) ([]string, error) {
	sqlCode := `SELECT bot_tokens FROM users WHERE email=$1;`
	var botTokens pq.StringArray = pq.StringArray{}
	row := repo.db.QueryRow(sqlCode, email)
	err := row.Scan(&botTokens)
	BotTokens := []string(botTokens)
	if err == nil {
		return BotTokens, nil
	} else if err == sql.ErrNoRows {
		return nil, errors.New("no email found")
	}
	return nil, err
}
func (repo *UserRepository) GetUserPassword(email string) (string, error) {
	sqlCode := `SELECT password FROM users WHERE email=$1;`
	var password string
	row := repo.db.QueryRow(sqlCode, email)
	err := row.Scan(&password)
	if err == nil {
		return password, nil
	} else if err == sql.ErrNoRows {
		return "", errors.New("no email found")
	}
	return "", err
}
func (repo *UserRepository) AddUserAuthToken(email string, token string) error {
	sqlCode := `
	UPDATE users
	SET auth_tokens = array_append(auth_tokens, $1)
	WHERE email = $2;`
	_, err := repo.db.Exec(sqlCode, token, email)
	return err
}

func (repo *UserRepository) AddUserBotToken(email string, token string) error {
	sqlCode := `
	UPDATE users
	SET bot_tokens = array_append(bot_tokens, $1)
	WHERE email = $2;`
	_, err := repo.db.Exec(sqlCode, token, email)
	return err
}

func (repo *UserRepository) DeleteUserAuthToken(email string, token string) error {
	sqlCode := `
	UPDATE users
	SET auth_tokens = ARRAY_REMOVE(auth_tokens, $1)
	WHERE email = $2;`
	_, err := repo.db.Exec(sqlCode, token, email)
	return err
}

func (repo *UserRepository) DeleteUserBotToken(email string, token string) error {
	sqlCode := `
	UPDATE users
	SET bot_tokens = ARRAY_REMOVE(bot_tokens, $1)
	WHERE email = $2;`
	_, err := repo.db.Exec(sqlCode, token, email)
	return err
}
