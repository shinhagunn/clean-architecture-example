package middlewares

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/shinhagunn/todo-backend/pkg/logger"
)

func GinLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path

		c.Next()

		cost := time.Since(start)
		logger.Info(
			1,
			"%d %s:%s Cost:%v",
			c.Writer.Status(),
			c.Request.Method,
			path,
			cost,
		)
	}
}
