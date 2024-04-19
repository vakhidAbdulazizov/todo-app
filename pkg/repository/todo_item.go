package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/vakhidAbdulazizov/todo-app/models"
	"strings"
)

type TodoItemDb struct {
	db *sqlx.DB
}

func NewTodoItemDb(db *sqlx.DB) *TodoItemDb {
	return &TodoItemDb{db: db}
}

func (d *TodoItemDb) Create(listId int, input models.TodoItem) (int, error) {
	tx, err := d.db.Begin()

	if err != nil {
		return 0, err
	}

	var itemId int

	createItemQuery := fmt.Sprintf(`INSERT INTO %s (title, description) values ($1, $2) RETURNING id`, todoItemsTable)

	row := tx.QueryRow(createItemQuery, input.Title, input.Description)

	err = row.Scan(&itemId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	createListItemQuery := fmt.Sprintf(`INSERT INTO %s (list_id, item_id) values ($1, $2)`, listsItemsTable)

	_, err = tx.Exec(createListItemQuery, listId, itemId)

	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return itemId, tx.Commit()
}

func (d *TodoItemDb) GetAll(userId int, listId int) ([]models.TodoItem, error) {
	var items []models.TodoItem

	query := fmt.Sprintf(`SELECT ti.id, ti.title, ti.description, ti.done FROM %s ti INNER JOIN %s li on li.item_id = ti.id
									INNER JOIN %s ul on ul.list_id = li.list_id WHERE li.list_id = $1 AND ul.user_id = $2`,
		todoItemsTable, listsItemsTable, usersListTable)

	if err := d.db.Select(&items, query, listId, userId); err != nil {
		return nil, err
	}
	return items, nil
}

func (d *TodoItemDb) GetById(userId int, itemId int) (models.TodoItem, error) {
	var item models.TodoItem

	query := fmt.Sprintf(`SELECT ti.id, ti.title, ti.description, ti.done FROM %s ti INNER JOIN %s li on li.item_id = ti.id
									INNER JOIN %s ul on ul.list_id = li.list_id WHERE ti.id = $1 AND ul.user_id = $2`,
		todoItemsTable, listsItemsTable, usersListTable)

	if err := d.db.Get(&item, query, itemId, userId); err != nil {
		return item, err
	}
	return item, nil
}

func (d *TodoItemDb) Delete(userId int, itemId int) error {
	query := fmt.Sprintf(`DELETE FROM %s ti USING %s li, %s ul
						WHERE ti.id = li.item_id AND li.list_id = ul.list_id AND ul.user_id =  $1 AND ti.id = $2`,
		todoItemsTable, listsItemsTable, usersListTable)

	_, err := d.db.Exec(query, userId, itemId)

	return err
}

func (d *TodoItemDb) Update(userId int, itemId int, input models.UpdateItemInput) error {
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

	if input.Done != nil {
		setValues = append(setValues, fmt.Sprintf("done=$%d", iter))
		args = append(args, *input.Done)
		iter++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf(`UPDATE %s ti SET %s FROM %s li, %s ul
			WHERE ti.id = li.item_id AND li.list_id = ul.list_id AND ul.user_id = $%d AND ti.id = $%d`,
		todoItemsTable, setQuery, listsItemsTable, usersListTable, iter, iter+1)

	args = append(args, userId, itemId)

	_, err := d.db.Exec(query, args...)

	return err
}
