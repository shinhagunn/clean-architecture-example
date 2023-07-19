package helper

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/shinhagunn/todo-backend/pkg/logger"
)

func (h Helper) BindAndValid(c *gin.Context, payload interface{}, prefix string) bool {
	var ve validator.ValidationErrors
	if err := c.ShouldBind(payload); err != nil {
		if errors.As(err, &ve) && len(ve) > 0 {
			logger.Error(2, getFirstErrorValid(prefix, ve))
			c.JSON(http.StatusBadRequest, getFirstErrorValid(prefix, ve))
			return false
		}

		logger.Error(2, err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return false
	}

	return true
}

func getFirstErrorValid(prefix string, ve validator.ValidationErrors) string {
	return fmt.Sprintf("%s.%s_%s", prefix, strings.ToLower(ve[0].Field()), ve[0].Tag())
}
