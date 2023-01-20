package router

import (
	"context"
	"github.com/chenshone/tiktok-lite/service"
	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	// test api
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/test", func(c *gin.Context) {
		service.GetName(1, context.Background())
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
}
