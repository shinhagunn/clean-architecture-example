package helper

import (
	"github.com/gin-gonic/gin"
	"github.com/shinhagunn/todo-backend/pkg/logger"
)

func (h Helper) ResponseError(c *gin.Context, code int, e error) {
	logger.Error(2, e.Error())
	c.JSON(code, e.Error())
}

func (h Helper) ResponseData(c *gin.Context, code int, data interface{}) {
	c.JSON(code, data)
}
