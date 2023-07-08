package usecases

import (
	"context"
	"testing"

	"github.com/shinhagunn/eug/filters"
	"github.com/shinhagunn/todo-backend/config"
	"github.com/shinhagunn/todo-backend/internal/models"
	"github.com/shinhagunn/todo-backend/pkg/postgres"
	"golang.org/x/crypto/bcrypt"
)

func TestUserUsecase(t *testing.T) {
	ctx := context.TODO()

	cfg, err := config.New()
	if err != nil {
		t.Fatalf("Failed to connect to config: %v", err)
	}

	db, err := postgres.New(cfg)
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}

	db.AutoMigrate(&models.User{})

	userUsecase := NewUserUsecase(db)

	newUser := &models.User{
		Email:    "test@gmail.com",
		Password: "12345678",
	}

	if err := userUsecase.Create(ctx, newUser); err != nil {
		t.Fatalf("Failed to create user: %v", err)
	}

	user, err := userUsecase.First(ctx, filters.WithFieldEqual("email", "test@gmail.com"))
	if err != nil {
		t.Fatalf("Failed to first user: %v", err)
	}

	if user.Email != "test@gmail.com" || !user.CheckPassword("12345678") {
		t.Fatalf("Data user save wrong: %v", err)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("00000000"), bcrypt.DefaultCost)
	if err != nil {
		t.Fatalf("Failed to hash password user: %v", err)
	}

	if err := userUsecase.Updates(ctx, user, &models.User{
		Password: string(hashedPassword),
	}); err != nil {
		t.Fatalf("Failed to update user: %v", err)
	}

	if user.Email != "test@gmail.com" || !user.CheckPassword("00000000") {
		t.Fatalf("Data user save wrong: %v", err)
	}

	if err := userUsecase.Delete(ctx, user); err != nil {
		t.Fatalf("Failed to delete user: %v", err)
	}

	user, err = userUsecase.First(ctx, filters.WithFieldEqual("email", "test@gmail.com"))
	if user != nil {
		t.Fatalf("Fail delete: %v", err)
	}
}
