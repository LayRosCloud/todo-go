package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	userTable = "users"
	todoListsTable = "todo_lists"
	usersListsTable ="users_lists"
	todoItemsTable = "todo_items"
	listsItemsTable = "lists_items"
)

type Config struct {
	Host string
	Port string
	Username string
	Password string
	SchemaName string
	SSLMode string
}

func NewPostgresDb(cfg Config) (*sqlx.DB, error) {
	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host,
		cfg.Port,
		cfg.Username,
		cfg.Password,
		cfg.SchemaName,
		cfg.SSLMode,	
	);
	db, err := sqlx.Open("postgres", connectionString)

	if err != nil {
		return nil, err
	}

	err = db.Ping();
	
	if err != nil {
		return nil, err
	}
	return db, err
}