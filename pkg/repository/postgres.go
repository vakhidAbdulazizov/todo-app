package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

const (
	userTable       = "users"
	todoListTable   = "todo_lists"
	usersListTable  = "users_list"
	todoItemsTable  = "todo_items"
	listsItemsTable = "lists_items"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgresDb(ctg Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		ctg.Host, ctg.Port, ctg.Username, ctg.DBName, ctg.Password, ctg.SSLMode))
	if err != nil {
		return nil, err
	}

	err = db.Ping()

	if err != nil {
		return nil, err
	}

	return db, nil
}
