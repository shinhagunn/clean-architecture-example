package middlewares

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shinhagunn/eug/filters"
	"github.com/shinhagunn/todo-backend/internal/usecases"
	"github.com/shinhagunn/todo-backend/pkg/jwt"
)

func Auth(userUsecase usecases.UserUsecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, "Authorization header missing")
			c.Abort()
			return
		}

		claims, err := jwt.ValidateToken(tokenString)
		if err != nil {
			fmt.Println(err.Error())
			c.JSON(http.StatusUnauthorized, err.Error())
			c.Abort()
			return
		}

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
