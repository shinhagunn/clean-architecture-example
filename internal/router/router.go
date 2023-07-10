package router

import (
	"github.com/gin-gonic/gin"
	"github.com/shinhagunn/todo-backend/internal/controllers/identity"
	"github.com/shinhagunn/todo-backend/internal/controllers/public"
	"github.com/shinhagunn/todo-backend/internal/controllers/resource"
	"github.com/shinhagunn/todo-backend/internal/usecases"
	"gorm.io/gorm"
)

func New(db *gorm.DB) {
	router := gin.New()

	userUsecase := usecases.NewUserUsecase(db)
	taskUsecase := usecases.NewTaskUsecase(db)

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	v2 := router.Group("/api/v2")

	public.NewRouter(v2)
	identity.NewRouter(v2, userUsecase)
	resource.NewRouter(v2, userUsecase, taskUsecase)

	router.Run(":3000")
}
