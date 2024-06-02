package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/layroscloud/todo-go/entity"
)

type Authorization interface {
	CreateUser(user entity.User) (int, error)
	FindByUsernameAndPassword(username, password string) (entity.User, error)
}

type TodoList interface {
	FindAllByUserId(userId int64) ([]entity.TodoList, error)
	FindById(id int64) (*entity.TodoList, error)
	Create(list entity.TodoList, userId int64) (int64, error)
	Update(list entity.TodoList) (int64, error)
	Delete(id int64) (int, error)
}

type TodoItem interface {
	//TODO
}

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		TodoList:      NewListPostgres(db),
	}
}
