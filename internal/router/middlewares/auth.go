package middlewares

import (
	"context"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/shinhagunn/eug/filters"
	"github.com/shinhagunn/todo-backend/internal/usecases"
	"github.com/shinhagunn/todo-backend/pkg/jwt"
)

func Auth(userUsecase usecases.UserUsecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" && !strings.Contains(tokenString, "Bearer") {
			c.JSON(http.StatusUnauthorized, "Authorization header missing")
			c.Abort()
			return
		}

		claims, err := jwt.ValidateToken(strings.TrimPrefix(tokenString, "Bearer "))
		if err != nil {
			c.JSON(http.StatusUnauthorized, err.Error())
			c.Abort()
			return
		}

		newTokenString, err := jwt.GenerateJWTToken(claims.UID)
		if err != nil {
			c.JSON(http.StatusUnauthorized, err.Error())
			c.Abort()
			return
		}

		c.Header("Authorization", "Bearer "+newTokenString)

		user, err := userUsecase.First(context.TODO(), filters.WithFieldEqual("uid", claims.UID))
		if err != nil {
			c.JSON(http.StatusUnauthorized, err.Error())
			c.Abort()
			return
		}

		c.Set("user", user)
		c.Next()
	}
}
