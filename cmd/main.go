package main

import (
	"log"

	"github.com/n1tro-dev/todo-my"
	"github.com/n1tro-dev/todo-my/pkg/handler"
	"github.com/n1tro-dev/todo-my/pkg/repository"
	"github.com/n1tro-dev/todo-my/pkg/service"
	"github.com/spf13/viper"
)

func main() {

	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	rep := repository.NewRepository()
	services := service.NewService(rep)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)

	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error with connect server: %s", err.Error())
	}

}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
