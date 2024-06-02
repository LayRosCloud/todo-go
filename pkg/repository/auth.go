package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/layroscloud/todo-go/entity"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{
		db: db,
	}
}

func (a AuthPostgres) CreateUser(user entity.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s(name, username, password) values ($1, $2, $3) RETURNING id", userTable)
	row := a.db.QueryRow(query, user.Name, user.Username, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (a AuthPostgres) FindByUsernameAndPassword(username, password string) (user entity.User, err error) {
	query := fmt.Sprintf("SELECT u.id FROM %s u WHERE u.username = $1 AND u.password = $2", userTable)
	err = a.db.Get(&user, query, username, password)
	return user, err
}
