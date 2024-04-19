package main

import (
	"context"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/vakhidAbdulazizov/todo-app"
	"github.com/vakhidAbdulazizov/todo-app/pkg/handler"
	"github.com/vakhidAbdulazizov/todo-app/pkg/repository"
	"github.com/vakhidAbdulazizov/todo-app/pkg/service"
	"os"
	"os/signal"
	"syscall"
)

// @title Todo App API
// @version 1.0
// @description API Server for TodoList Application

// @host localhost:8003
// @BasePath /

// @securityDefinitions.apikey JWTTokenAuth
// @in header
// @name Authorization
func main() {

	logrus.SetFormatter(new(logrus.JSONFormatter))
	//if err := initConfig(); err != nil {
	//	logrus.Fatalf("error initialization confid: %s", err.Error())
	//}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error load env file: %s", err.Error())
	}

	db, err := repository.NewPostgresDb(repository.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Username: os.Getenv("DB_USER_NAME"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	})

	if err != nil {
		logrus.Fatalf("error connect db: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	srv := new(todo.Server)
	go func() {
		err = srv.Run(os.Getenv("APP_PORT"), handlers.InitRoutes())
		if err != nil {
			logrus.Fatalf("error run server: %s", err.Error())
		}
	}()

	logrus.Print("TodoApp started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	logrus.Print("TodoApp Shutting down")
	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occurred on server stutting down: %s", err.Error())
	}

	if err := db.Close(); err != nil {
		logrus.Errorf("error occurred on db connection close: %s", err.Error())
	}
}

/*
func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}

*/
