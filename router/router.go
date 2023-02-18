package router

import (
	"github.com/chenshone/tiktok-lite/controller"
	"github.com/chenshone/tiktok-lite/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
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
	feed.GET("/", middleware.JWTAuth(), controller.GetVideoListByLastTime)

	// favorite api 点赞操作
	favorite := r.Group("/douyin/favorite/")
	favorite.GET("list/", middleware.JWTAuth(), controller.GetFavoriteList)
	favorite.POST("action/", middleware.JWTAuth(), controller.AddOrCancelFavorite)

	// comment api
	comment := r.Group("/douyin/comment/")
	comment.POST("action/", middleware.JWTAuth(), controller.CommentAction)
	comment.GET("list/", middleware.JWTAuth(), controller.GetCommentList)

	// relation api 关注操作
	relation := r.Group("/douyin/relation/")
	relation.POST("action/", middleware.JWTAuth(), controller.FollowUserOrNot)
	relation.GET("follow/list/", middleware.JWTAuth(), controller.GetFollowList)
	relation.GET("follower/list/", middleware.JWTAuth(), controller.GetFollowerList)
	relation.GET("friend/list/", middleware.JWTAuth(), controller.GetFriendList)

	// message api
	message := r.Group("/douyin/message/")
	message.POST("action/", middleware.JWTAuth(), controller.SendMessage)
	message.GET("chat/", middleware.JWTAuth(), controller.GetMessageList)
}
