package service

import (
	"context"
	"github.com/chenshone/tiktok-lite/dal/query"
)

type UserInfo struct {
	ID            int    `json:"id"`
	Username      string `json:"username"`
	Avatar        string `json:"avatar"`
	FollowCount   int    `json:"follow_count"`
	FollowerCount int    `json:"follower_count"`
}

func GetUser(id int) (*UserInfo, error) {
	user := query.Q.User
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
		Avatar:        data[0].Avatar,
		FollowCount:   int(data[0].FollowCount),
		FollowerCount: int(data[0].FollowerCount),
	}, nil
}
