package service

import (
	"github.com/vakhidAbdulazizov/todo-app/models"
	"github.com/vakhidAbdulazizov/todo-app/pkg/repository"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
	GenerateToken(username string, password string) (string, error)
	ParseToken(token string) (int, error)
	RestorePassword(email string, confirmKey string, password string) error
	ForgotPassword(email string) error
}

type TodoList interface {
	Create(userId int, list models.TodoList) (int, error)
	GetAll(userId int) ([]models.TodoList, error)
	GetByID(userId int, listId int) (models.TodoList, error)
	Delete(userId int, listId int) error
	Update(userId int, listId int, input models.UpdateListInput) error
}

type TodoItem interface {
	Create(userId int, listId int, input models.TodoItem) (int, error)
	GetAll(userId int, listId int) ([]models.TodoItem, error)
	GetById(userId int, itemId int) (models.TodoItem, error)
	Delete(userId int, itemId int) error
	Update(userId int, itemId int, input models.UpdateItemInput) error
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repo.Authorization),
		TodoList:      NewTodoListService(repo.TodoList),
		TodoItem:      NewTodoItemService(repo.TodoItem, repo.TodoList),
	}
}
