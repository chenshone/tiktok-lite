package service

import (
	"context"
	"github.com/chenshone/tiktok-lite/dal/model"
	"time"
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

func GetVideoListByUserId(userId int) ([]*VideoInfo, error) {
	video := q.Video
	vdo := video.WithContext(context.Background())
	data, err := vdo.Where(video.UserID.Eq(int32(userId))).Preload(video.Author).Find()
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
			Author: Author{
				ID:            int(v.Author.ID),
				Username:      v.Author.Username,
				FollowCount:   int(v.Author.FollowCount),
				FollowerCount: int(v.Author.FollowerCount),
				IsFollow:      false,
			},
			IsFavorite: false,
		})
	}
	return videos, nil
}

// GetVideoListByTime 按发布时间倒序获取最新的视频列表
func GetVideoListByTime(lastTime time.Time) ([]*VideoInfo, error) {
	v := q.Video
	vdo := v.WithContext(context.Background())
	data, err := vdo.Where(v.CreateAt.Lt(lastTime)).Preload(v.Author).Order(v.CreateAt.Desc()).Limit(30).Find()
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
			Author: Author{
				ID:            int(v.Author.ID),
				Username:      v.Author.Username,
				FollowCount:   int(v.Author.FollowCount),
				FollowerCount: int(v.Author.FollowerCount),
				IsFollow:      false,
			},
			IsFavorite: false,
		})
	}
	return videos, nil
}
