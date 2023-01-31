package router

import (
	"github.com/chenshone/tiktok-lite/controller"
	"github.com/gin-gonic/gin"
)

type loginOrRegisterReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func InitRouter(r *gin.Engine) {
	// test api
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// user api
	user := r.Group("/douyin/user")
	user.GET("/", func(c *gin.Context) {
		userID := c.Query("user_id")
		//token := c.Query("token")
		userInfo := controller.GetUserInfo(userID)
		c.JSON(200, userInfo)
	})
	user.POST("/login", func(c *gin.Context) {
		userInfo := loginOrRegisterReq{}
		_ = c.ShouldBindJSON(&userInfo)
		data := controller.Login(userInfo.Username, userInfo.Password)
		c.JSON(200, data)
	})
	user.POST("/register", func(c *gin.Context) {
		userInfo := loginOrRegisterReq{}
		_ = c.ShouldBindJSON(&userInfo)
		data := controller.Register(userInfo.Username, userInfo.Password)
		c.JSON(200, data)
	})
}
