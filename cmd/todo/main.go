package main

import (
	"github.com/shinhagunn/todo-backend/config"
	"github.com/shinhagunn/todo-backend/internal/models"
	"github.com/shinhagunn/todo-backend/internal/router"
	"github.com/shinhagunn/todo-backend/pkg/postgres"
)

func main() {
	if err := config.New(); err != nil {
		panic(err)
	}

	db, err := postgres.New(config.Cfg)
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&models.User{}, &models.Task{}, &models.Category{})

	router.New(db)
}
