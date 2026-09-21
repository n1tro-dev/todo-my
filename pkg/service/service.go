package service

import (
	"github.com/n1tro-dev/todo-my"
	"github.com/n1tro-dev/todo-my/pkg/repository"
)

type Authorizaton interface {
	CreateUser(user todo.User) (int, error)
	GenerateToken(username, password string) (string, error)
}

type TodoList interface {
}

type TodoItem interface {
}

type Service struct {
	Authorizaton
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorizaton: NewAuthServise(repos.Authorizaton),
	}
}
