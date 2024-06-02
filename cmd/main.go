package main

import (
	"os"

	"github.com/joho/godotenv"
	todoapp "github.com/layroscloud/todo-go"
	"github.com/layroscloud/todo-go/pkg/handler"
	"github.com/layroscloud/todo-go/pkg/repository"
	"github.com/layroscloud/todo-go/pkg/service"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	if err := InitConfig(); err != nil {
		logrus.Fatalf("error config: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error config: %s", err.Error())
	}

	db, err := repository.NewPostgresDb(repository.Config{
		Host:       viper.GetString("datasource.host"),
		Port:       viper.GetString("datasource.port"),
		Username:   viper.GetString("datasource.username"),
		SSLMode:    viper.GetString("datasource.sslmode"),
		SchemaName: viper.GetString("datasource.schema"),
		Password:   os.Getenv("DB_PASSWORD"),
	})

	if err != nil {
		logrus.Fatalf("database initialize configuration: %s", err.Error())
	}

	server := new(todoapp.Server)
	repositories := repository.NewRepository(db)
	services := service.NewService(repositories)
	handlers := handler.NewHandler(services)

	if err := server.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		logrus.Fatalf("error: %s", err.Error())
	}
}

func InitConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
