package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/vakhidAbdulazizov/todo-app/models"
	"strings"
)

type TodoListDb struct {
	db *sqlx.DB
}

func NewTodoListDb(db *sqlx.DB) *TodoListDb {
	return &TodoListDb{db: db}
}

func (r *TodoListDb) Create(userId int, list models.TodoList) (int, error) {
	tx, err := r.db.Begin()

	if err != nil {
		return 0, err
	}

	var id int

	createListQuery := fmt.Sprintf("INSERT INTO %s (title, description) VALUES ($1, $2) RETURNING id", todoListTable)
	row := tx.QueryRow(createListQuery, list.Title, list.Description)

	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	createUsersListQuery := fmt.Sprintf("INSERT INTO %s (user_id, list_id) VALUES ($1, $2)", usersListTable)

	_, err = tx.Exec(createUsersListQuery, userId, id)

	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()
}

func (r *TodoListDb) GetAll(userId int) ([]models.TodoList, error) {
	var lists []models.TodoList

	query := fmt.Sprintf("SELECT tl.id, tl.title, tl.description FROM %s tl INNER JOIN %s ul on tl.id=ul.list_id WHERE ul.user_id = $1",
		todoListTable, usersListTable)

	err := r.db.Select(&lists, query, userId)

	return lists, err
}

func (r *TodoListDb) GetByID(userId int, listId int) (models.TodoList, error) {
	var list models.TodoList

	query := fmt.Sprintf("SELECT tl.id, tl.title, tl.description FROM %s tl INNER JOIN %s ul on tl.id=ul.list_id WHERE ul.user_id = $1 AND ul.list_id = $2",
		todoListTable, usersListTable)

	err := r.db.Get(&list, query, userId, listId)

	return list, err
}

func (r *TodoListDb) Delete(userId int, listId int) error {
	query := fmt.Sprintf("DELETE FROM %s tl USING %s ul WHERE tl.id = ul.list_id AND ul.user_id=$1 AND ul.list_id=$2",
		todoListTable, usersListTable)

	_, err := r.db.Exec(query, userId, listId)

	return err
}

func (r *TodoListDb) Update(userId int, listId int, input models.UpdateListInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	iter := 1

	if input.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", iter))
		args = append(args, *input.Title)
		iter++
	}

	if input.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", iter))
		args = append(args, *input.Description)
		iter++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s tl SET %s FROM %s ul WHERE tl.id = ul.list_id AND ul.list_id=$%d AND ul.user_id=$%d",
		todoListTable, setQuery, usersListTable, iter, iter+1)

	args = append(args, listId, userId)

	_, err := r.db.Exec(query, args...)

	return err
}
