package helper

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
)

func (h Helper) BindAndValid(c *gin.Context, payload interface{}) error {
	var ve validator.ValidationErrors
	if err := c.ShouldBind(payload); err != nil {
		if errors.As(err, &ve) && len(ve) > 0 {
			// logger.Error(2, getFirstErrorValid(ve))
			return errors.New(getFirstErrorValid(ve))
		}

		// logger.Error(2, err.Error())
		return err
	}

	return nil
}

func getFirstErrorValid(ve validator.ValidationErrors) string {
	return fmt.Sprintf("%s_%s", strings.ToLower(ve[0].Field()), ve[0].Tag())
}
