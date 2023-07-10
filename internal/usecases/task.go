package usecases

import (
	"github.com/shinhagunn/eug"
	"github.com/shinhagunn/todo-backend/internal/models"
	"gorm.io/gorm"
)

type taskUsecase struct {
	eug.Usecase[models.Task]
}

type TaskUsecase interface {
	eug.IUsecase[models.Task]
}

func NewTaskUsecase(db *gorm.DB) TaskUsecase {
	return taskUsecase{
		Usecase: eug.Usecase[models.Task]{
			Repo: eug.NewRepository[models.Task](db, models.Task{}),
		},
	}
}
