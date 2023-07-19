package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shinhagunn/todo-backend/pkg/logger"
)

func GinRecover() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				logger.Error(4, "[Recover from panic] Error: %v", err)

				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()

		c.Next()
	}
}
