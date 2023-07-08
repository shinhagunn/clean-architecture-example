package usecases

import (
	"github.com/shinhagunn/eug"
	"github.com/shinhagunn/todo-backend/internal/models"
	"gorm.io/gorm"
)

type userUsecase struct {
	eug.Usecase[models.User]
}

type UserUsecase interface {
	eug.IUsecase[models.User]
}

func NewUserUsecase(db *gorm.DB) UserUsecase {
	return userUsecase{
		Usecase: eug.Usecase[models.User]{
			Repo: eug.NewRepository[models.User](db, models.User{}),
		},
	}
}
