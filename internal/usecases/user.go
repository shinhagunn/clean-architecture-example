package usecases

import (
	"context"
	"errors"

	"github.com/shinhagunn/eug"
	"github.com/shinhagunn/todo-backend/internal/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	First(context.Context, *models.User, ...eug.Filter) error
	Find(context.Context, []models.User, ...eug.Filter) error
	Create(context.Context, *models.User) error
	Updates(context.Context, *models.User, *models.User) error
	Delete(context.Context, *models.User) error
}

type userUsecase struct {
	repo UserRepository
}

func NewUserUsecase(repo UserRepository) *userUsecase {
	return &userUsecase{repo}
}

func (u *userUsecase) First(ctx context.Context, filters ...eug.Filter) (*models.User, error) {
	user := &models.User{}

	err := u.repo.First(ctx, user, filters...)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *userUsecase) Find(ctx context.Context, filters ...eug.Filter) (*models.User, error) {
	user := &models.User{}

	if err := u.repo.First(ctx, user, filters...); err != nil {
		return nil, err
	}

	return user, nil
}

func (u *userUsecase) Create(ctx context.Context, user *models.User) error {
	return u.repo.Create(ctx, user)
}

func (u *userUsecase) Updates(ctx context.Context, user *models.User, updates *models.User) error {
	return u.repo.Updates(ctx, user, updates)
}

func (u *userUsecase) Delete(ctx context.Context, user *models.User) error {
	return u.repo.Delete(ctx, user)
}
