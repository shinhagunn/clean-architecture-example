package helpers

import (
	"errors"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type APIError struct {
	Code    int
	Message string
}

type Helpers struct{}

func (h Helpers) ParserData(c *gin.Context, payload interface{}, prefix string) error {
	var ve validator.ValidationErrors
	if err := c.ShouldBind(payload); err != nil {
		if errors.As(err, &ve) && len(ve) > 0 {
			return errors.New(fmt.Sprintf("%s.%s_%s", prefix, strings.ToLower(ve[0].Field()), ve[0].Tag()))
		}

		return err
	}

	return nil
}
