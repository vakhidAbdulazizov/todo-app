package service

import (
	"github.com/vakhidAbdulazizov/todo-app/models"
	"github.com/vakhidAbdulazizov/todo-app/pkg/repository"
)

type TodoItemService struct {
	repo     repository.TodoItem
	listRepo repository.TodoList
}

func NewTodoItemService(repo repository.TodoItem, listRepo repository.TodoList) *TodoItemService {
	return &TodoItemService{repo: repo, listRepo: listRepo}
}

func (s *TodoItemService) Create(userId int, listId int, input models.TodoItem) (int, error) {
	_, err := s.listRepo.GetByID(userId, listId)

	if err != nil {
		return 0, err
	}

	return s.repo.Create(listId, input)
}

func (s *TodoItemService) GetAll(userId int, listId int) ([]models.TodoItem, error) {
	return s.repo.GetAll(userId, listId)
}

func (s *TodoItemService) GetById(userId int, itemId int) (models.TodoItem, error) {
	return s.repo.GetById(userId, itemId)
}

func (s *TodoItemService) Delete(userId int, itemId int) error {
	return s.repo.Delete(userId, itemId)
}

func (s *TodoItemService) Update(userId int, itemId int, input models.UpdateItemInput) error {
	return s.repo.Update(userId, itemId, input)
}
