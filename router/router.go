package router

import (
	"github.com/chenshone/tiktok-lite/controller"
	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	// test api
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// user api
	r.GET("/user/:id", func(c *gin.Context) {
		userId := c.Param("id")
		userInfo := controller.GetUser(userId)
		c.JSON(userInfo.Code, userInfo)
	})
}
