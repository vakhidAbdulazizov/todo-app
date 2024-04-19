package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/vakhidAbdulazizov/todo-app/models"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
	GetUser(username string, password string) (models.User, error)
	RestorePassword(email string, confirmKey string, password string) error
	ForgotPassword(email string, hashKey string, confirmKey string) error
}

type TodoList interface {
	Create(userId int, list models.TodoList) (int, error)
	GetAll(userId int) ([]models.TodoList, error)
	GetByID(userId int, listId int) (models.TodoList, error)
	Delete(userId int, listId int) error
	Update(userId int, listId int, input models.UpdateListInput) error
}

type TodoItem interface {
	Create(listId int, input models.TodoItem) (int, error)
	GetAll(userId int, listId int) ([]models.TodoItem, error)
	GetById(userId int, itemId int) (models.TodoItem, error)
	Delete(userId int, itemId int) error
	Update(userId int, itemId int, input models.UpdateItemInput) error
}

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthDb(db),
		TodoList:      NewTodoListDb(db),
		TodoItem:      NewTodoItemDb(db),
	}
}
