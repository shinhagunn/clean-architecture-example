package identity

import (
	"github.com/gin-gonic/gin"
	"github.com/shinhagunn/todo-backend/internal/helper"
	"github.com/shinhagunn/todo-backend/internal/usecases"
)

type Handler struct {
	helper.Helper
	userUsecase usecases.UserUsecase
}

func NewRouter(
	router *gin.RouterGroup,
	userUsecase usecases.UserUsecase,
) {
	handler := &Handler{
		userUsecase: userUsecase,
	}

	identity := router.Group("/identity")

	identity.POST("/register", handler.Register)
	identity.POST("/login", handler.Login)
}
