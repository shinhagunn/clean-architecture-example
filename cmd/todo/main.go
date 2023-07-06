package main

import (
	"log"

	"github.com/shinhagunn/todo-backend/config"
	"github.com/shinhagunn/todo-backend/pkg/postgres"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		panic(err)
	}

	db, err := postgres.New(cfg)
	if err != nil {
		panic(err)
	}

	// userUsecase := usecases.NewUserUsecase(repositories.NewUserRepository(db))
	// userUsecase.Create(&model.User{
	// 	Email:    "test@gmail.com",
	// 	Password: "12345678",

	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(userUsecase)

	log.Fatal(db)
}
