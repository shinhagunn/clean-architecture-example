package middlewares

import (
	"context"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
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
			help.ResponseError(c, http.StatusUnauthorized, errors.New("authorization header missing"))
			c.Abort()
			return
		}

		claims, err := util.ParseToken(strings.TrimPrefix(tokenString, "Bearer "))
		if err != nil {
			help.ResponseError(c, http.StatusUnauthorized, errors.Wrap(err, "parse token fail"))
			c.Abort()
			return
		}

		newTokenString, err := util.GenerateToken(claims.UID)
		if err != nil {
			help.ResponseError(c, http.StatusUnauthorized, errors.Wrap(err, "generate token fail"))
			c.Abort()
			return
		}

		c.Header("Authorization", "Bearer "+newTokenString)

		user, err := userUsecase.First(context.TODO(), filters.WithFieldEqual("uid", claims.UID))
		if err != nil {
			help.ResponseError(c, http.StatusUnauthorized, errors.Wrap(err, "user not found"))
			c.Abort()
			return
		}

		c.Set("user", user)
		c.Next()
	}
}
