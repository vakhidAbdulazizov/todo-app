package service

import (
	"github.com/vakhidAbdulazizov/todo-app/models"
	"github.com/vakhidAbdulazizov/todo-app/pkg/repository"
)

type TodoListService struct {
	repo repository.TodoList
}

func NewTodoListService(repo repository.TodoList) *TodoListService {
	return &TodoListService{repo: repo}
}

func (s *TodoListService) Create(userId int, list models.TodoList) (int, error) {
	return s.repo.Create(userId, list)
}

func (s *TodoListService) GetAll(userId int) ([]models.TodoList, error) {
	return s.repo.GetAll(userId)
}

func (s *TodoListService) GetByID(userId int, listId int) (models.TodoList, error) {
	return s.repo.GetByID(userId, listId)
}

func (s *TodoListService) Delete(userId int, listId int) error {
	return s.repo.Delete(userId, listId)
}

func (s *TodoListService) Update(userId int, listId int, input models.UpdateListInput) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return s.repo.Update(userId, listId, input)
}
