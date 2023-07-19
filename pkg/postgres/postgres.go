package postgres

import (
	"fmt"

	"github.com/shinhagunn/todo-backend/internal/models"
	"github.com/shinhagunn/todo-backend/pkg/setting"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Setup() *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai", setting.Cfg.DB.Host, setting.Cfg.DB.User, setting.Cfg.DB.Password, setting.Cfg.DB.Name, setting.Cfg.DB.Port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		// Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&models.User{}, &models.Task{}, &models.Category{})

	return db
}
