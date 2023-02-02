package service

import (
	"context"
	"errors"
	"github.com/chenshone/tiktok-lite/dal/model"
	"github.com/chenshone/tiktok-lite/util/util"
	"strconv"
	"strings"
)

type UserInfo struct {
	ID            int    `json:"id"`
	Username      string `json:"name"`
	FollowCount   int    `json:"follow_count"`
	FollowerCount int    `json:"follower_count"`
	IsFollow      bool   `json:"is_follow"`
}

type UserToken struct {
	ID    int
	Token string
}

func GetUserInfo(id int) (*UserInfo, error) {
	user := q.User
	userDo := user.WithContext(context.Background())
	data, err := userDo.Where(user.ID.Eq(int32(id))).Find()
	if err != nil {
		return nil, err
	}
	if len(data) == 0 {
		return &UserInfo{}, nil
	}
	return &UserInfo{
		ID:            int(data[0].ID),
		Username:      data[0].Username,
		FollowCount:   int(data[0].FollowCount),
		FollowerCount: int(data[0].FollowerCount),
		// TODO: 是否关注
		IsFollow: false,
	}, nil
}

func Register(username string, password string) error {
	user := q.User
	do := user.WithContext(context.Background())
	username = strings.Trim(username, " ")
	if len(username) == 0 {
		return errors.New("用户名不能为空")
	}
	if len(password) == 0 {
		return errors.New("密码不能为空")
	}
	if len(username) > 32 {
		return errors.New("用户名长度不能超过32")
	}
	if len(password) > 32 {
		return errors.New("密码长度不能超过32")
	}
	data, err := do.Where(user.Username.Eq(username)).Find()
	if err != nil {
		return err
	}
	if len(data) != 0 {
		return errors.New("用户名已存在")
	}

	newUser := model.User{
		Username: username,
		Password: password,
		Avatar: "https://pics6.baidu.com/feed/ca1349540923dd546f63018e776583d89d8248b2." +
			"jpeg?token=9573aa8647e48fce0ad1814fc07ce257&s=50B0AD7341D077E9492984CC0300F0E3",
	}
	err = do.Create(&newUser)
	return err
}

func Login(username string, password string) (*UserToken, error) {
	user := q.User
	do := user.WithContext(context.Background())
	username = strings.Trim(username, " ")
	if len(username) == 0 {
		return nil, errors.New("用户名不能为空")
	}
	if len(password) == 0 {
		return nil, errors.New("密码不能为空")
	}
	if len(username) > 32 {
		return nil, errors.New("用户名长度不能超过32")
	}
	if len(password) > 32 {
		return nil, errors.New("密码长度不能超过32")
	}
	data, err := do.Where(user.Username.Eq(username)).Find()
	if err != nil {
		return nil, err
	}
	if len(data) == 0 {
		return nil, errors.New("用户名/密码错误")
	}
	if data[0].Password != password {
		return nil, errors.New("用户名/密码错误")
	}

	jwt := util.JWT{}
	token := jwt.GenerateToken(strconv.Itoa(int(data[0].ID)), 1)
	return &UserToken{
		ID:    int(data[0].ID),
		Token: token,
	}, nil
}
