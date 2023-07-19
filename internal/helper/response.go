package helper

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/shinhagunn/todo-backend/pkg/logger"
)

type APIError struct {
	Code int   `json:"code"`
	Err  error `json:"msg"`
}

func NewAPIError(code int, msg string) APIError {
	return APIError{Code: code, Err: errors.New(msg)}
}

func (h Helper) ResponseError(c *gin.Context, e APIError) {
	logger.Error(2, e.Err.Error())
	c.JSON(e.Code, e.Err.Error())
}

func (h Helper) ResponseData(c *gin.Context, code int, data interface{}) {
	c.JSON(code, data)
}
