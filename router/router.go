package router

import (
	"github.com/chenshone/tiktok-lite/controller"
	"github.com/chenshone/tiktok-lite/middleware"
	"github.com/chenshone/tiktok-lite/util/util"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"strconv"
	"time"
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
	user.GET("/", middleware.JWTAuth(), func(c *gin.Context) {
		userID := c.Query("user_id")
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

	// publish api
	publish := r.Group("/douyin/publish")
	publish.POST("/action", middleware.JWTAuth(), func(c *gin.Context) {
		userID := c.GetInt("user_id")
		title := c.PostForm("title")
		video, err := c.FormFile("data")
		if err != nil {
			c.JSON(200, &gin.H{
				"status_code": -1,
				"status_msg":  "视频获取失败",
			})
			return
		}
		if video.Size > 10*1024*1024 {
			c.JSON(200, &gin.H{
				"status_code": -1,
				"status_msg":  "视频过大",
			})
			return
		}
		if !util.CheckExt(video.Filename) {
			c.JSON(200, &gin.H{
				"status_code": -1,
				"status_msg":  "视频格式错误",
			})
			return
		}
		log.Println("upload video: ", video.Filename)
		filename := strconv.Itoa(int(time.Now().Unix())) + "---" + video.Filename
		folderPath := util.Mkdir("./assets/video/")
		videoPath := folderPath + filename
		log.Println("视频上传: ", videoPath)
		err = c.SaveUploadedFile(video, videoPath)
		if err != nil {
			c.JSON(200, &gin.H{
				"status_code": -1,
				"status_msg":  "视频上传失败",
			})
			return
		}
		coverPath := videoPath + ".png"
		err = util.GetVideoCover(videoPath, coverPath, 1)
		if err != nil {
			_ = os.Remove(videoPath)
			c.JSON(200, &gin.H{
				"status_code": -1,
				"status_msg":  "视频上传失败",
			})
			return
		}
		log.Println("上传视频成功")
		data := controller.PublishVideo(userID, videoPath, coverPath, title)
		c.JSON(200, data)
	})
	publish.GET("/list", middleware.JWTAuth(), func(c *gin.Context) {
		userID := c.GetInt("user_id")
		data := controller.GetVideoList(userID)
		c.JSON(200, data)
	})
}
