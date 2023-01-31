package controller

import (
	"github.com/chenshone/tiktok-lite/service"
	"strconv"
)

type UserInfoRes struct {
	Code int         `json:"status_code"`
	Msg  string      `json:"status_msg"`
	User interface{} `json:"user"`
}

func GetUserInfo(id string) *UserInfoRes {
	userId, err := strconv.Atoi(id)
	if err != nil {
		return &UserInfoRes{
			Code: -1,
			Msg:  "无效id",
			User: nil,
		}
	}
	userInfo, err := service.GetUserInfo(userId)
	if err != nil {
		return &UserInfoRes{
			Code: -1,
			Msg:  err.Error(),
			User: nil,
		}
	}
	return &UserInfoRes{
		Code: 0,
		Msg:  "success",
		User: userInfo,
	}
}

type RegisterAndLoginRes struct {
	Code   int    `json:"status_code"`
	Msg    string `json:"status_msg"`
	UserID int    `json:"user_id"`
	Token  string `json:"token"`
}

func Register(username string, password string) *RegisterAndLoginRes {
	err := service.Register(username, password)
	if err != nil {
		return &RegisterAndLoginRes{
			Code: -1,
			Msg:  err.Error(),
		}
	}
	data, err := service.Login(username, password)
	if err != nil {
		return &RegisterAndLoginRes{
			Code: -1,
			Msg:  err.Error(),
		}
	}

	return &RegisterAndLoginRes{
		Code:   0,
		Msg:    "success",
		UserID: data.ID,
		Token:  data.Token,
	}
}

func Login(username string, password string) *RegisterAndLoginRes {
	data, err := service.Login(username, password)
	if err != nil {
		return &RegisterAndLoginRes{
			Code: -1,
			Msg:  err.Error(),
		}
	}

	return &RegisterAndLoginRes{
		Code:   0,
		Msg:    "success",
		UserID: data.ID,
		Token:  data.Token,
	}
}
