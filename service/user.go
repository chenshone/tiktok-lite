package service

import (
	"context"
	"errors"
	"github.com/chenshone/tiktok-lite/conf"
	"github.com/chenshone/tiktok-lite/dal/model"
	"github.com/chenshone/tiktok-lite/util/util"
	"strconv"
	"strings"
)

type UserInfo struct {
	ID              int    `json:"id"`
	Username        string `json:"name"`
	FollowCount     int    `json:"follow_count"`
	FollowerCount   int    `json:"follower_count"`
	IsFollow        bool   `json:"is_follow"`
	Avatar          string `json:"avatar"`
	BackgroundImage string `json:"background_image"`
	Signature       string `json:"signature"`
	TotalFavorited  int    `json:"total_favorited"`
	WorkCount       int    `json:"work_count"`
	FavoriteCount   int    `json:"favorite_count"`
}

type UserToken struct {
	ID    int
	Token string
}

func GetUserInfo(userID, targetUserId int) (*UserInfo, error) {
	user := q.User
	userDo := user.WithContext(context.Background())
	data, err := userDo.Where(user.ID.Eq(int32(targetUserId))).Find()
	if err != nil {
		return nil, err
	}
	if len(data) == 0 {
		return &UserInfo{}, nil
	}
	r := q.Relation
	rdo := r.WithContext(context.Background())
	resp, err := rdo.Where(r.UserID.Eq(int32(userID)), r.ToUserID.Eq(int32(targetUserId))).Find()
	if err != nil {
		return nil, err
	}
	return &UserInfo{
		ID:              int(data[0].ID),
		Username:        data[0].Username,
		FollowCount:     int(data[0].FollowCount),
		FollowerCount:   int(data[0].FollowerCount),
		IsFollow:        len(resp) > 0,
		Avatar:          data[0].Avatar,
		Signature:       data[0].Signature,
		BackgroundImage: data[0].BackgroundImage,
		TotalFavorited:  int(data[0].TotalFavorited),
		WorkCount:       int(data[0].WorkCount),
		FavoriteCount:   int(data[0].FavoriteCount),
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
		Username:        username,
		Password:        util.Md5Salt(password),
		Avatar:          conf.BaseURL + "assets/avator.png",
		BackgroundImage: conf.BaseURL + "assets/bg.jpeg",
		Signature:       "hello",
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
	if data[0].Password != util.Md5Salt(password) {
		return nil, errors.New("用户名/密码错误")
	}

	jwt := util.JWT{}
	token := jwt.GenerateToken(strconv.Itoa(int(data[0].ID)), 1)
	return &UserToken{
		ID:    int(data[0].ID),
		Token: token,
	}, nil
}
