package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/layroscloud/todo-go/entity"
)

type ListPostgres struct {
	db *sqlx.DB
}

func NewListPostgres(db *sqlx.DB) *ListPostgres {
	return &ListPostgres{
		db: db,
	}
}

func (l *ListPostgres) FindAllByUserId(userId int64) ([]entity.TodoList, error) {
	var items []entity.TodoList
	query := fmt.Sprintf("SELECT tl.id as id, tl.title as title, tl.description as description FROM %s tl INNER JOIN %s ul ON ul.list_id = tl.id WHERE ul.user_id=$1", todoListsTable, usersListsTable)
	err := l.db.Select(&items, query, userId)
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (l *ListPostgres) FindById(id int64) (*entity.TodoList, error) {
	var item entity.TodoList
	query := fmt.Sprintf("SELECT tl.id, tl.title, tl.description FROM %s tl WHERE tl.id=$1", todoListsTable)
	err := l.db.Select(&item, query, id)
	if err != nil {
		return nil, err
	}
	return &item, err
}

func (l *ListPostgres) Create(list entity.TodoList, userId int64) (int64, error) {
	var id int64
	var tempId int64
	query := fmt.Sprintf("INSERT INTO %s (title, description) VALUES ($1, $2) RETURNING id", todoListsTable)
	row := l.db.QueryRow(query, list.Title, list.Description)
	err := row.Scan(&id)
	if err != nil {
		return 0, err
	}

	query = fmt.Sprintf("INSERT INTO %s (list_id, user_id) VALUES ($1, $2) RETURNING id", usersListsTable)
	row = l.db.QueryRow(query, id, userId)
	err = row.Scan(&tempId)
	if err != nil {
		return 0, err
	}
	return id, err
}

func (l *ListPostgres) Update(list entity.TodoList) (int64, error) {
	query := fmt.Sprintf("UPDATE %s SET title=$1, description=$2 WHERE id=$3", todoListsTable)
	l.db.QueryRow(query, list.Title, list.Description, list.Id)
	return list.Id, nil
}

func (l *ListPostgres) Delete(id int64) (int, error) {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$3", todoListsTable)
	l.db.QueryRow(query, id)
	return 1, nil
}
