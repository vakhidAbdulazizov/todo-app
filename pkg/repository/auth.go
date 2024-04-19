package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/vakhidAbdulazizov/todo-app/models"
	"github.com/vakhidAbdulazizov/todo-app/pkg/mail"
)

type AuthDb struct {
	db *sqlx.DB
}

func NewAuthDb(db *sqlx.DB) *AuthDb {
	return &AuthDb{db: db}
}

func (a *AuthDb) CreateUser(user models.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, username, password_hash, email) values($1,$2,$3, $4) RETURNING id", userTable)
	row := a.db.QueryRow(query, user.Name, user.UserName, user.Password, user.Email)

	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (a *AuthDb) GetUser(username string, password string) (models.User, error) {
	var user models.User

	query := fmt.Sprintf("SELECT id FROM %s WHERE username=$1 AND password_hash=$2", userTable)

	err := a.db.Get(&user, query, username, password)

	return user, err
}

func (a *AuthDb) RestorePassword(email string, confirmKey string, password string) error {
	var user models.User

	query := fmt.Sprintf("SELECT id FROM %s WHERE email=$1 AND confirm_hash=$2", userTable)

	err := a.db.Get(&user, query, email, confirmKey)

	if err != nil {
		return err
	}

	setNewPasswordQuery := fmt.Sprintf("UPDATE %s SET password_hash=$1 WHERE id=$2", userTable)

	_, err = a.db.Exec(setNewPasswordQuery, password, user.Id)

	return err
}

func (a *AuthDb) ForgotPassword(email string, hashKey string, confirmKey string) error {
	var user models.User

	query := fmt.Sprintf("SELECT id, email FROM %s WHERE email=$1", userTable)

	err := a.db.Get(&user, query, email)

	if err != nil {
		return err
	}

	err = mail.SendEmail(email, []byte("Confirm code: \n"+confirmKey))

	if err != nil {
		return err
	}

	setConfirmKeyQuery := fmt.Sprintf("UPDATE %s SET confirm_hash=$1 WHERE id=$2", userTable)

	_, err = a.db.Exec(setConfirmKeyQuery, hashKey, user.Id)

	return err
}
