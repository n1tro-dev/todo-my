package service

import (
	"crypto/sha1"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/n1tro-dev/todo-my"
	"github.com/n1tro-dev/todo-my/pkg/repository"
)

const (
	salt       = "qwrqr123123qweqwe"
	tokenTTL   = 12 * time.Hour
	signingKey = "qwerwt#23442QWRETQWRsdfsdfasdf"
)

type tokenClaims struct {
	jwt.StandardClaims
	User_id int `json:"user_id"`
}

type AuthServise struct {
	repo repository.Authorization
}

func NewAuthServise(repo repository.Authorization) *AuthServise {
	return &AuthServise{repo: repo}
}

func (s *AuthServise) CreateUser(user todo.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func (s *AuthServise) GenerateToken(username, password string) (string, error) {

	user, err := s.repo.GetUser(username, generatePasswordHash(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.ID,
	})

	return token.SignedString([]byte(signingKey))
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
