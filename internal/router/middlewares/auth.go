package middlewares

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/shinhagunn/eug/filters"
	"github.com/shinhagunn/todo-backend/internal/helper"
	"github.com/shinhagunn/todo-backend/internal/usecases"
	"github.com/shinhagunn/todo-backend/pkg/util"
)

func Auth(userUsecase usecases.UserUsecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		help := helper.Helper{}

		tokenString := c.GetHeader("Authorization")
		if tokenString == "" && !strings.Contains(tokenString, "Bearer") {
			help.ResponseError(c, helper.APIError{Code: http.StatusUnauthorized, Err: errors.New("Authorization header missing")})
			c.Abort()
			return
		}

		claims, err := util.ParseToken(strings.TrimPrefix(tokenString, "Bearer "))
		if err != nil {
			help.ResponseError(c, helper.APIError{Code: http.StatusUnauthorized, Err: err})
			c.Abort()
			return
		}

		newTokenString, err := util.GenerateToken(claims.UID)
		if err != nil {
			help.ResponseError(c, helper.APIError{Code: http.StatusUnauthorized, Err: err})
			c.Abort()
			return
		}

		c.Header("Authorization", "Bearer "+newTokenString)

		user, err := userUsecase.First(context.TODO(), filters.WithFieldEqual("uid", claims.UID))
		if err != nil {
			help.ResponseError(c, helper.APIError{Code: http.StatusUnauthorized, Err: err})
			c.Abort()
			return
		}

		c.Set("user", user)
		c.Next()
	}
}
