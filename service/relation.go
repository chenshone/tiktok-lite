package service

import (
	"context"
	"github.com/chenshone/tiktok-lite/dal/model"
)

func FollowUser(userID, toUserID int) error {
	r := q.Relation
	rdo := r.WithContext(context.Background())
	data := &model.Relation{
		UserID:   int32(userID),
		ToUserID: int32(toUserID),
	}
	resp, err := rdo.Where(r.UserID.Eq(int32(userID)), r.ToUserID.Eq(int32(toUserID))).Find()
	if err != nil {
		return err
	}

	if len(resp) > 0 {
		return nil
	}

	err = rdo.Create(data)
	if err != nil {
		return err
	}
	return nil
}

func UnFollowUser(userID, toUserID int) error {
	r := q.Relation
	rdo := r.WithContext(context.Background())
	_, err := rdo.Where(r.UserID.Eq(int32(userID)), r.ToUserID.Eq(int32(toUserID))).Delete()
	if err != nil {
		return err
	}
	return nil
}
