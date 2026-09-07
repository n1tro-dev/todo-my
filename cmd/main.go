package main

import (
	"log"

	"github.com/n1tro-dev/todo-my"
	"github.com/n1tro-dev/todo-my/pkg/handler"
)

func main() {

	handlers := new(handler.Handler)

	srv := new(todo.Server)

	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("error with connect server: %s", err.Error())
	}

}
