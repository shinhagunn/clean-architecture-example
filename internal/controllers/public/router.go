package public

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter(router *gin.RouterGroup) {
	public := router.Group("/public")

	public.GET("/test", func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	})
}
