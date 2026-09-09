package repository

import (
	"github.com/jmoiron/sqlx"
)

type Authorizaton interface {
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
	return &Repository{}
}
