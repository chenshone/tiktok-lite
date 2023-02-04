package controller

import (
	"errors"
	"github.com/chenshone/tiktok-lite/service"
	"time"
)

type PublishVideoResp struct {
	Code int    `json:"status_code"`
	Msg  string `json:"status_msg"`
}

func PublishVideo(userId int, videoPath string, coverPath string, title string) *PublishVideoResp {
	err := service.PublishVideo(userId, videoPath, coverPath, title)
	if err != nil {
		return &PublishVideoResp{
			Code: -1,
			Msg:  err.Error(),
		}
	}
	return &PublishVideoResp{
		Code: 0,
		Msg:  "success",
	}
}

type VideoListResp struct {
	Code      int         `json:"status_code"`
	Msg       string      `json:"status_msg"`
	VideoList interface{} `json:"video_list"`
}

func GetVideoListByUserId(userIf int) *VideoListResp {
	videoList, err := service.GetVideoListByUserId(userIf)
	if err != nil {
		return &VideoListResp{
			Code: -1,
			Msg:  err.Error(),
		}
	}
	return &VideoListResp{
		Code:      0,
		Msg:       "success",
		VideoList: videoList,
	}
}

func GetVideoListByLastTime(lastTime string) *VideoListResp {
	t, err := time.Parse("2006-01-02 15:04:05", lastTime)
	if err != nil {
		return &VideoListResp{
			Code: -1,
			Msg:  errors.New("时间格式错误").Error(),
		}
	}
	videoList, err := service.GetVideoListByTime(t)
	if err != nil {
		return &VideoListResp{
			Code: -1,
			Msg:  err.Error(),
		}
	}
	return &VideoListResp{
		Code:      0,
		Msg:       "success",
		VideoList: videoList,
	}
}
