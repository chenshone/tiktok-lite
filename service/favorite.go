package service

import (
	"context"
	"github.com/chenshone/tiktok-lite/dal"
	"github.com/chenshone/tiktok-lite/dal/model"
	"github.com/chenshone/tiktok-lite/dal/query"
	"log"
)

func GetFavoriteList(userID, targetUserID int) ([]*videoInfo, error) {
	fav := q.Favorite
	v := q.Video
	favDo := fav.WithContext(context.Background())
	vdo := v.WithContext(context.Background())

	favoriteIDs, err := favDo.Select(fav.VideoID).Where(fav.UserID.Eq(int32(targetUserID))).Find()
	if err != nil {
		return nil, err
	}
	ids := make([]int32, len(favoriteIDs))
	for i, v := range favoriteIDs {
		ids[i] = v.VideoID
	}
	data, err := vdo.Where(v.ID.In(ids...)).Preload(v.Author).Find()
	if err != nil {
		return nil, err
	}
	videoList := make([]*videoInfo, len(data))

	r := q.Relation
	rdo := r.WithContext(context.Background())
	var isFollow, isFavorite bool

	for i, v := range data {
		isFavorite, isFollow = false, false
		if _, err := rdo.Where(r.UserID.Eq(int32(userID)), r.ToUserID.Eq(v.Author.ID)).First(); err == nil {
			isFollow = true
		}
		if _, err := favDo.Where(fav.UserID.Eq(int32(userID)), fav.VideoID.Eq(v.ID)).First(); err == nil {
			isFavorite = true
		}
		videoList[i] = &videoInfo{
			ID:            int(v.ID),
			PlayURL:       v.PlayURL,
			CoverURL:      v.CoverURL,
			FavoriteCount: int(v.FavoriteCount),
			CommentCount:  int(v.CommentCount),
			Title:         v.Title,
			IsFavorite:    isFavorite,
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
		}
	}
	return videoList, nil
}

// AddOrCancelFavorite
/*
	IsAdd: 1 -> add
	IsAdd: 2 -> cancel
*/
func AddOrCancelFavorite(userID, videoID, IsAdd int) (err error) {
	tq := query.Use(dal.DB)
	tx := tq.Begin()
	defer func() {
		if recover() != nil || err != nil {
			_ = tx.Rollback()
		}
	}()
	fav := tx.Favorite
	favDo := fav.WithContext(context.Background())
	v := tx.Video
	vdo := v.WithContext(context.Background())
	u := tx.User
	udo := u.WithContext(context.Background())

	videoAuthor, err := vdo.Select(v.UserID).Where(v.ID.Eq(int32(videoID))).First()
	if err != nil {
		return err
	}

	if IsAdd == 1 { // add
		if _, err = favDo.Where(fav.UserID.Eq(int32(userID)), fav.VideoID.Eq(int32(videoID))).First(); err == nil {
			return err
		}
		newFav := model.Favorite{
			UserID:  int32(userID),
			VideoID: int32(videoID),
		}
		if err = favDo.Create(&newFav); err != nil {
			return err
		}
		if _, err = vdo.Where(v.ID.Eq(int32(videoID))).Update(v.FavoriteCount, v.FavoriteCount.Add(1)); err != nil {
			return err
		}
		if _, err = udo.Where(u.ID.Eq(videoAuthor.UserID)).Update(u.TotalFavorited,
			u.TotalFavorited.Add(1)); err != nil {
			return err
		}
		if _, err = udo.Where(u.ID.Eq(int32(userID))).Update(u.FavoriteCount, u.FavoriteCount.Add(1)); err != nil {
			return err
		}
		return tx.Commit()
	}
	log.Printf("取消点赞\n")
	res, err := favDo.Where(fav.UserID.Eq(int32(userID)), fav.VideoID.Eq(int32(videoID))).Delete()
	if err != nil {
		return err
	}
	if res.RowsAffected > 0 {
		if _, err = vdo.Where(v.ID.Eq(int32(videoID))).Update(v.FavoriteCount,
			v.FavoriteCount.Sub(1)); err != nil {
			return err
		}
	}
	if _, err = udo.Where(u.ID.Eq(videoAuthor.UserID)).Update(u.TotalFavorited,
		u.TotalFavorited.Sub(1)); err != nil {
		return err
	}
	if _, err = udo.Where(u.ID.Eq(int32(userID))).Update(u.FavoriteCount, u.FavoriteCount.Sub(1)); err != nil {
		return err
	}
	return tx.Commit()
}
