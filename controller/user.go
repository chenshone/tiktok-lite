package controller

import (
	"github.com/chenshone/tiktok-lite/service"
	"strconv"
)

func GetUser(id string) *PageData {
	userId, err := strconv.Atoi(id)
	if err != nil {
		return &PageData{
			Code: 404,
			Msg:  "无效id",
		}
	}
	userInfo, err := service.GetUser(userId)
	if err != nil {
		return &PageData{
			Code: 500,
			Msg:  err.Error(),
		}
	}
	return &PageData{
		Code: 200,
		Msg:  "success",
		Data: userInfo,
	}
}
