package adapters

import (
	"database/sql"
	"fmt"
)

type PostgressDb struct {
	Host     string
	Port     string
	Username string
	Password string
	Dbname   string
	PsqlInfo string
}

func NewPostgressDb(host string, port string, username string, password string, dbname string) *PostgressDb {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+"password=%s dbname=%s sslmode=disable", host, port, username, password, dbname)
	return &PostgressDb{Host: host, Port: port, Username: username, Password: password, Dbname: dbname, PsqlInfo: psqlInfo}
}

func (p *PostgressDb) Connect() (*sql.DB, error) {
	db, err := sql.Open("postgres", p.PsqlInfo)
	return db, err
}

func (p *PostgressDb) Close(db *sql.DB) {
	db.Close()
}
