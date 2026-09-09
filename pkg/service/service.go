package service

import "github.com/n1tro-dev/todo-my/pkg/repository"

type Authorizaton interface {

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

func NewService(repos *repository.Repository) *Service{
	return &Service{}
}