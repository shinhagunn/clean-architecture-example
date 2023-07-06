package repositories

import (
	"context"

	"github.com/shinhagunn/eug"
	"github.com/shinhagunn/todo-backend/internal/models"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type userRepository struct {
	db *gorm.DB
	schema.Tabler
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db, &models.User{}}
}

func (r *userRepository) First(ctx context.Context, user *models.User, filters ...eug.Filter) error {
	return eug.ApplyFilters(r.db.WithContext(ctx).Table(r.TableName()), filters).First(&user).Error
}

func (r *userRepository) Find(ctx context.Context, users []models.User, filters ...eug.Filter) error {
	return eug.ApplyFilters(r.db.WithContext(ctx).Table(r.TableName()), filters).Find(&users).Error
}

func (r *userRepository) Create(ctx context.Context, user *models.User) error {
	return r.db.Create(&user).Error
}

func (r *userRepository) Updates(ctx context.Context, user *models.User, updates *models.User) error {
	return r.db.Model(&user).Updates(&updates).Error
}

func (r *userRepository) Delete(ctx context.Context, user *models.User) error {
	return r.db.Delete(&user).Error
}
