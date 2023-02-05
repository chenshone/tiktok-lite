package controller

import (
	"errors"
	"github.com/chenshone/tiktok-lite/service"
	"github.com/chenshone/tiktok-lite/util/util"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"strconv"
	"time"
)

type PublishVideoResp struct {
	Code int    `json:"status_code"`
	Msg  string `json:"status_msg"`
}

func PublishVideo(c *gin.Context) {
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
	err = service.PublishVideo(userID, videoPath, coverPath, title)
	if err != nil {
		c.JSON(200, &PublishVideoResp{
			Code: -1,
			Msg:  err.Error(),
		})
		return
	}
	c.JSON(200, &PublishVideoResp{
		Code: 0,
		Msg:  "success",
	})
}

type VideoListResp struct {
	Code      int         `json:"status_code"`
	Msg       string      `json:"status_msg"`
	VideoList interface{} `json:"video_list"`
}

func GetVideoListByUserId(c *gin.Context) {
	userID := c.GetInt("user_id")
	videoList, err := service.GetVideoListByUserId(userID)
	if err != nil {
		c.JSON(200, &VideoListResp{
			Code: -1,
			Msg:  err.Error(),
		})
		return
	}
	c.JSON(200, &VideoListResp{
		Code:      0,
		Msg:       "success",
		VideoList: videoList,
	})
}

func GetVideoListByLastTime(c *gin.Context) {
	lastTime := c.Query("last_time")
	_ = c.Query("token")
	if lastTime == "" {
		lastTime = time.Now().Format("2006-01-02 15:04:05")
	}
	t, err := time.Parse("2006-01-02 15:04:05", lastTime)
	if err != nil {
		c.JSON(200, &VideoListResp{
			Code: -1,
			Msg:  errors.New("时间格式错误").Error(),
		})
		return
	}
	videoList, err := service.GetVideoListByTime(t)
	if err != nil {
		c.JSON(200, &VideoListResp{
			Code: -1,
			Msg:  err.Error(),
		})
		return
	}
	c.JSON(200, &VideoListResp{
		Code:      0,
		Msg:       "success",
		VideoList: videoList,
	})
}
