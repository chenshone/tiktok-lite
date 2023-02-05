package service

import (
	"context"
	"github.com/chenshone/tiktok-lite/dal/model"
)

func GetFavoriteList(userID int) ([]*VideoInfo, error) {
	fav := q.Favorite
	v := q.Video
	favDo := fav.WithContext(context.Background())
	vdo := v.WithContext(context.Background())

	favoriteIDs, err := favDo.Select(fav.VideoID).Where(fav.UserID.Eq(int32(userID))).Find()
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
	videoList := make([]*VideoInfo, len(data))
	for i, v := range data {
		videoList[i] = &VideoInfo{
			ID:            int(v.ID),
			PlayURL:       v.PlayURL,
			CoverURL:      v.CoverURL,
			FavoriteCount: int(v.FavoriteCount),
			CommentCount:  int(v.CommentCount),
			Title:         v.Title,
			IsFavorite:    true,
			Author: Author{
				ID:            int(v.Author.ID),
				Username:      v.Author.Username,
				FollowCount:   int(v.Author.FollowCount),
				FollowerCount: int(v.Author.FollowerCount),
				IsFollow:      false,
			},
		}
	}
	return videoList, nil
}

// AddOrCancelFavorite
/*
	IsAdd: true -> add
	IsAdd: false -> cancel
*/
func AddOrCancelFavorite(userID, videoID, IsAdd int) error {
	fav := q.Favorite
	favDo := fav.WithContext(context.Background())

	if IsAdd == 1 {
		newFav := model.Favorite{
			UserID:  int32(userID),
			VideoID: int32(videoID),
		}
		err := favDo.Create(&newFav)
		return err
	}
	_, err := favDo.Where(fav.UserID.Eq(int32(userID)), fav.VideoID.Eq(int32(videoID))).Delete()

	return err
}
