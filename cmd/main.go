package main

import (
	"log"

	"github.com/n1tro-dev/todo-my"
	"github.com/n1tro-dev/todo-my/pkg/handler"
	"github.com/n1tro-dev/todo-my/pkg/repository"
	"github.com/n1tro-dev/todo-my/pkg/service"
)

func main() {

	rep := repository.NewRepository()
	services := service.NewService(rep)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)

	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("error with connect server: %s", err.Error())
	}

}
