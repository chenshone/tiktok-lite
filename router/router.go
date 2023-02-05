package router

import (
	"github.com/chenshone/tiktok-lite/controller"
	"github.com/chenshone/tiktok-lite/middleware"
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
	user := r.Group("/douyin/user/")
	user.GET("/", middleware.JWTAuth(), controller.GetUserInfo)
	user.POST("login/", controller.Login)
	user.POST("register/", controller.Register)

	// publish api
	publish := r.Group("/douyin/publish/")
	publish.POST("action/", middleware.JWTAuth(), controller.PublishVideo)
	publish.GET("list/", middleware.JWTAuth(), controller.GetVideoListByUserId)

	// feed api
	feed := r.Group("/douyin/feed/")
	feed.GET("/", controller.GetVideoListByLastTime)

	// favorite api
	favorite := r.Group("/douyin/favorite/")
	favorite.GET("list/", middleware.JWTAuth(), controller.GetFavoriteList)
	favorite.POST("action/", middleware.JWTAuth(), controller.AddOrCancelFavorite)
}
