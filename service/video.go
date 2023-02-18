package service

import (
	"context"
	"github.com/chenshone/tiktok-lite/conf"
	"github.com/chenshone/tiktok-lite/dal"
	"github.com/chenshone/tiktok-lite/dal/model"
	"github.com/chenshone/tiktok-lite/dal/query"
	"time"
)

func PublishVideo(userId int, videoPath string, coverPath string, title string) (err error) {
	tq := query.Use(dal.DB)
	tx := tq.Begin()
	defer func() {
		if recover() != nil || err != nil {
			_ = tx.Rollback()
		}
	}()
	video := tx.Video
	do := video.WithContext(context.Background())

	newVedio := model.Video{
		UserID:   int32(userId),
		PlayURL:  conf.BaseURL + videoPath,
		CoverURL: conf.BaseURL + coverPath,
		Title:    title,
	}
	err = do.Create(&newVedio)
	if err != nil {
		return err
	}

	u := tx.User
	udo := u.WithContext(context.Background())
	_, err = udo.Where(u.ID.Eq(int32(userId))).Update(u.WorkCount, u.WorkCount.Add(1))
	if err != nil {
		return err
	}
	return tx.Commit()
}

type videoInfo struct {
	ID            int    `json:"id"`
	Author        author `json:"author"`
	PlayURL       string `json:"play_url"`
	CoverURL      string `json:"cover_url"`
	FavoriteCount int    `json:"favorite_count"`
	CommentCount  int    `json:"comment_count"`
	IsFavorite    bool   `json:"is_favorite"`
	Title         string `json:"title"`
}
type author struct {
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

func GetVideoListByUserId(userID, targetUserID int) ([]*videoInfo, error) {
	video := q.Video
	vdo := video.WithContext(context.Background())
	data, err := vdo.Where(video.UserID.Eq(int32(targetUserID))).Preload(video.Author).Find()
	if err != nil {
		return nil, err
	}

	// 是否关注该用户
	var isFollow bool
	r := q.Relation
	rdo := r.WithContext(context.Background())
	if _, err := rdo.Where(r.UserID.Eq(int32(userID)), r.ToUserID.Eq(int32(targetUserID))).First(); err != nil {
		isFollow = false
	} else {
		isFollow = true
	}

	var isFavorite bool
	f := q.Favorite
	fdo := f.WithContext(context.Background())
	videos := make([]*videoInfo, len(data))
	for i, v := range data {
		if _, err := fdo.Where(f.UserID.Eq(int32(userID)), f.VideoID.Eq(v.ID)).First(); err != nil {
			isFavorite = false
		} else {
			isFavorite = true
		}

		videos[i] = &videoInfo{
			ID:            int(v.ID),
			PlayURL:       v.PlayURL,
			CoverURL:      v.CoverURL,
			FavoriteCount: int(v.FavoriteCount),
			CommentCount:  int(v.CommentCount),
			Title:         v.Title,
			Author: author{
				ID:              int(v.Author.ID),
				Username:        v.Author.Username,
				FollowCount:     int(v.Author.FollowCount),
				FollowerCount:   int(v.Author.FollowerCount),
				IsFollow:        isFollow,
				Avatar:          v.Author.Avatar,
				BackgroundImage: v.Author.BackgroundImage,
				Signature:       v.Author.Signature,
				TotalFavorited:  int(v.Author.TotalFavorited),
				WorkCount:       int(v.Author.WorkCount),
				FavoriteCount:   int(v.Author.FavoriteCount),
			},
			IsFavorite: isFavorite,
		}
	}
	return videos, nil
}

// GetVideoListByTime 按发布时间倒序获取最新的视频列表
func GetVideoListByTime(lastTime time.Time, userID int) ([]*videoInfo, error) {
	v := q.Video
	vdo := v.WithContext(context.Background())
	data, err := vdo.Where(v.CreateAt.Lt(lastTime)).Preload(v.Author).Order(v.CreateAt.Desc()).Limit(30).Find()
	if err != nil {
		return nil, err
	}

	var isFollow bool
	var isFavorite bool
	f := q.Favorite
	fdo := f.WithContext(context.Background())
	r := q.Relation
	rdo := r.WithContext(context.Background())
	videos := make([]*videoInfo, len(data))
	for i, v := range data {
		isFollow, isFavorite = false, false
		if userID != -1 { // -1 表示未登录
			// 是否关注该用户
			if _, err := rdo.Where(r.UserID.Eq(int32(userID)), r.ToUserID.Eq(v.UserID)).First(); err == nil {
				isFollow = true
			}
			// 是否收藏该视频
			if _, err := fdo.Where(f.UserID.Eq(int32(userID)), f.VideoID.Eq(v.ID)).First(); err == nil {
				isFavorite = true
			}
		}
		videos[i] = &videoInfo{
			ID:            int(v.ID),
			PlayURL:       v.PlayURL,
			CoverURL:      v.CoverURL,
			FavoriteCount: int(v.FavoriteCount),
			CommentCount:  int(v.CommentCount),
			Title:         v.Title,
			Author: author{
				ID:              int(v.Author.ID),
				Username:        v.Author.Username,
				FollowCount:     int(v.Author.FollowCount),
				FollowerCount:   int(v.Author.FollowerCount),
				IsFollow:        isFollow,
				Avatar:          v.Author.Avatar,
				BackgroundImage: v.Author.BackgroundImage,
				Signature:       v.Author.Signature,
				TotalFavorited:  int(v.Author.TotalFavorited),
				WorkCount:       int(v.Author.WorkCount),
				FavoriteCount:   int(v.Author.FavoriteCount),
			},
			IsFavorite: isFavorite,
		}
	}
	return videos, nil
}
