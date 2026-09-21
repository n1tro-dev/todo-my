package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/n1tro-dev/todo-my"
)

type Authorizaton interface {
	CreateUser(user todo.User) (int, error)
}

type TodoList interface {
}

type TodoItem interface {
}

type Repository struct {
	Authorizaton
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorizaton: NewAuthPostgres(db),
	}
}
