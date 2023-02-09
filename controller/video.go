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

type publishVideoResp struct {
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
	if video.Size > 100*1024*1024 { // 100M
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
	// 获取视频封面并上传文件
	coverPath := videoPath + ".png"
	err = util.GetVideoCover(videoPath, coverPath, 1)
	if err != nil {
		c.JSON(200, &gin.H{
			"status_code": -1,
			"status_msg":  "视频上传失败",
		})
		return
	}
	// 上传视频
	err = c.SaveUploadedFile(video, videoPath)
	if err != nil {
		// 视频上传失败，删除封面
		_ = os.Remove(coverPath)
		c.JSON(200, &gin.H{
			"status_code": -1,
			"status_msg":  "视频上传失败",
		})
		return
	}
	log.Println("上传视频成功")
	err = service.PublishVideo(userID, videoPath, coverPath, title)
	if err != nil {
		//数据库存储失败，删除视频和封面
		_ = os.Remove(videoPath)
		_ = os.Remove(coverPath)
		c.JSON(200, &publishVideoResp{
			Code: -1,
			Msg:  err.Error(),
		})
		return
	}
	c.JSON(200, &publishVideoResp{
		Code: 0,
		Msg:  "success",
	})
}

type videoListResp struct {
	Code      int         `json:"status_code"`
	Msg       string      `json:"status_msg"`
	VideoList interface{} `json:"video_list"`
}

func GetVideoListByUserId(c *gin.Context) {
	userID := c.GetInt("user_id")
	query, ok := c.GetQuery("user_id")
	if !ok {
		c.JSON(200, &videoListResp{
			Code: -1,
			Msg:  errors.New("参数缺失").Error(),
		})
		return
	}
	targetUserID, err := strconv.Atoi(query)
	if err != nil {
		c.JSON(200, &videoListResp{
			Code: -1,
			Msg:  errors.New("参数错误").Error(),
		})
		return
	}
	videoList, err := service.GetVideoListByUserId(userID, targetUserID)
	if err != nil {
		c.JSON(200, &videoListResp{
			Code: -1,
			Msg:  err.Error(),
		})
		return
	}
	c.JSON(200, &videoListResp{
		Code:      0,
		Msg:       "success",
		VideoList: videoList,
	})
}

func GetVideoListByLastTime(c *gin.Context) {
	lastTime := c.Query("last_time")
	if lastTime == "" {
		lastTime = time.Now().Format("2006-01-02 15:04:05")
	}
	t, err := time.Parse("2006-01-02 15:04:05", lastTime)
	if err != nil {
		c.JSON(200, &videoListResp{
			Code: -1,
			Msg:  errors.New("时间格式错误").Error(),
		})
		return
	}
	userID := -1
	_, ok := c.GetQuery("token")
	if ok {
		userID = c.GetInt("user_id")
	}
	videoList, err := service.GetVideoListByTime(t, userID)
	if err != nil {
		c.JSON(200, &videoListResp{
			Code: -1,
			Msg:  err.Error(),
		})
		return
	}
	c.JSON(200, &videoListResp{
		Code:      0,
		Msg:       "success",
		VideoList: videoList,
	})
}
