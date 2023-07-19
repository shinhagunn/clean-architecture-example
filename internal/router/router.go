package router

import (
	"github.com/gin-gonic/gin"
	"github.com/shinhagunn/todo-backend/internal/controllers/identity"
	"github.com/shinhagunn/todo-backend/internal/controllers/public"
	"github.com/shinhagunn/todo-backend/internal/controllers/resource"
	"github.com/shinhagunn/todo-backend/internal/router/middlewares"
	"github.com/shinhagunn/todo-backend/internal/usecases"
	"gorm.io/gorm"
)

func InitRouter(db *gorm.DB) *gin.Engine {
	router := gin.New()
	router.Use(middlewares.GinLogger(), middlewares.GinRecover())

	userUsecase := usecases.NewUserUsecase(db)
	taskUsecase := usecases.NewTaskUsecase(db)

	v2 := router.Group("/api/v2")

	public.NewRouter(v2)
	identity.NewRouter(v2, userUsecase)
	resource.NewRouter(v2, userUsecase, taskUsecase)

	return router
}
