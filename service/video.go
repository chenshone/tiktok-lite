package service

import (
	"context"
	"errors"
	"github.com/chenshone/tiktok-lite/dal/model"
)

func PublishVideo(userId int, videoPath string, coverPath string, title string) error {
	video := q.Video
	do := video.WithContext(context.Background())

	newVedio := model.Video{
		UserID:   int32(userId),
		PlayURL:  videoPath,
		CoverURL: coverPath,
		Title:    title,
	}
	err := do.Create(&newVedio)
	return err
}

type VideoInfo struct {
	ID            int    `json:"id"`
	Author        Author `json:"author"`
	PlayURL       string `json:"play_url"`
	CoverURL      string `json:"cover_url"`
	FavoriteCount int    `json:"favorite_count"`
	CommentCount  int    `json:"comment_count"`
	IsFavorite    bool   `json:"is_favorite"`
	Title         string `json:"title"`
}
type Author struct {
	ID            int    `json:"id"`
	Username      string `json:"name"`
	FollowCount   int    `json:"follow_count"`
	FollowerCount int    `json:"follower_count"`
	IsFollow      bool   `json:"is_follow"`
}

func GetVideoList(userId int) ([]*VideoInfo, error) {
	video := q.Video
	u := q.User
	udo := u.WithContext(context.Background())
	vdo := video.WithContext(context.Background())

	var user Author
	err := udo.Where(u.ID.Eq(int32(userId))).Scan(&user)
	if err != nil {
		return nil, errors.New("用户不存在")
	}

	data, err := vdo.Where(video.UserID.Eq(int32(userId))).Find()
	if err != nil {
		return nil, err
	}

	var videos []*VideoInfo
	for _, v := range data {
		videos = append(videos, &VideoInfo{
			ID:            int(v.ID),
			PlayURL:       v.PlayURL,
			CoverURL:      v.CoverURL,
			FavoriteCount: int(v.FavoriteCount),
			CommentCount:  int(v.CommentCount),
			Title:         v.Title,
			Author:        user,
			IsFavorite:    false,
		})
	}
	return videos, nil
}
