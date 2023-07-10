package resource

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shinhagunn/todo-backend/internal/helpers"
	"github.com/shinhagunn/todo-backend/internal/router/middlewares"
	"github.com/shinhagunn/todo-backend/internal/usecases"
)

type Handler struct {
	helpers.Helpers
	userUsecase usecases.UserUsecase
	taskUsecase usecases.TaskUsecase
}

func NewRouter(
	router *gin.RouterGroup,
	userUsecase usecases.UserUsecase,
	taskUsecase usecases.TaskUsecase,
) {
	handler := &Handler{
		userUsecase: userUsecase,
		taskUsecase: taskUsecase,
	}

	resource := router.Group("/resource")
	resource.Use(middlewares.Auth(userUsecase))
	resource.GET("/test", func(c *gin.Context) { c.String(http.StatusOK, "OK") })

	// /tasks
	resource.GET("/tasks", handler.GetTasks)
	resource.POST("/tasks", handler.CreateTask)
	resource.PUT("/tasks", handler.UpdateTask)
	resource.DELETE("/tasks", handler.DeleteTask)

	// /categories
}
