package main

import (
	"fmt"
	"os"

	app "github.com/ShawaDev/auth"
	"github.com/ShawaDev/auth/pkg/handler"
	"github.com/ShawaDev/auth/pkg/repository"
	"github.com/ShawaDev/auth/pkg/service"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

func main() {
	server := new(app.Server)

	if err := godotenv.Load(); err != nil {
		fmt.Print(err)
	}

	if err := InitConfig(); err != nil {
		fmt.Print(err)
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.user"),
		Password: os.Getenv("password"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})

	if err != nil {
		fmt.Print(err)
	}

	repo := repository.NewRepository(db)
	service := service.NewService(repo)
	handler := handler.NewHandler(service)

	if err := server.Run(viper.GetString("port"), handler.InitRoutes()); err != nil {
		fmt.Print(err)
	}
}

func InitConfig() error {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
