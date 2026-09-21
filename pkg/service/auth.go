package service

import (
	"crypto/sha1"
	"fmt"

	"github.com/n1tro-dev/todo-my"
	"github.com/n1tro-dev/todo-my/pkg/repository"
)

const salt = "qwrqr123123qweqwe"

type AuthService struct {
	repo repository.Authorizaton
}

func NewAuthServise(repo repository.Authorizaton) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user todo.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
