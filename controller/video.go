package controller

import "github.com/chenshone/tiktok-lite/service"

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

func GetVideoList(userIf int) *VideoListResp {
	videoList, err := service.GetVideoList(userIf)
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
