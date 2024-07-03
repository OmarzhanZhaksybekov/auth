package main

import (
	"os"

	app "github.com/ShawaDev/auth"
	"github.com/ShawaDev/auth/pkg/handler"
	"github.com/ShawaDev/auth/pkg/repository"
	"github.com/ShawaDev/auth/pkg/service"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	server := new(app.Server)
	logrus.SetFormatter(new(logrus.JSONFormatter))

	//Enabling env varibales
	if err := godotenv.Load(); err != nil {
		logrus.Fatal(err)
	}

	//Enabling config
	if err := InitConfig(); err != nil {
		logrus.Fatal(err)
	}

	//Connecting to DB
	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.user"),
		Password: os.Getenv("password"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})

	if err != nil {
		logrus.Fatal(err)
	}

	repo := repository.NewRepository(db)
	service := service.NewService(repo)
	handler := handler.NewHandler(service)

	//Running server on port from config
	if err := server.Run(viper.GetString("port"), handler.InitRoutes()); err != nil {
		logrus.Fatal(err)
	}
}

// Enables to read config
func InitConfig() error {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
