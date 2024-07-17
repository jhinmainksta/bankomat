package main

import (
	"log"

	"github.com/jhinmainksta/bankomat"
	"github.com/jhinmainksta/bankomat/pkg/handler"
	"github.com/jhinmainksta/bankomat/pkg/service"
	"github.com/spf13/viper"
)

func main() {

	if err := InitConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	accounts := []service.BankAccount{}
	handlers := handler.NewHandler(accounts)

	srv := new(bankomat.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatal("error occured while running http server: %s", err.Error())

	}
}

func InitConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
