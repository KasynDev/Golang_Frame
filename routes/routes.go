package routes

import (
	"go_frame/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetUp() *gin.Engine {
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	r.GET("/", func(c *gin.Context) {
		c.JSONP(http.StatusOK, gin.H{
			"message": "Welcome Gin Server",
		})
	})
	return r
}
