package usecases

import (
	"github.com/shinhagunn/eug"
	"github.com/shinhagunn/todo-backend/internal/models"
	"gorm.io/gorm"
)

type categoryUsecase struct {
	eug.Usecase[models.Category]
}

type CategoryUsecase interface {
	eug.IUsecase[models.Category]
}

func NewCategoryUsecase(db *gorm.DB) CategoryUsecase {
	return categoryUsecase{
		Usecase: eug.Usecase[models.Category]{
			Repo: eug.NewRepository[models.Category](db, models.Category{}),
		},
	}
}
