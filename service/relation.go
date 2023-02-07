package service

import (
	"context"
	"github.com/chenshone/tiktok-lite/dal"
	"github.com/chenshone/tiktok-lite/dal/model"
	"github.com/chenshone/tiktok-lite/dal/query"
)

func FollowUser(userID, toUserID int) (err error) {
	tq := query.Use(dal.DB)
	tx := tq.Begin() // 开启事务
	defer func() {
		if recover() != nil || err != nil {
			_ = tx.Rollback()
		}
	}()

	r := tx.Relation
	rdo := r.WithContext(context.Background())
	data := &model.Relation{
		UserID:   int32(userID),
		ToUserID: int32(toUserID),
	}
	//防止重复关注
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
	u := tx.User
	udo := u.WithContext(context.Background())
	if _, err = udo.Where(u.ID.Eq(int32(toUserID))).Update(u.FollowerCount, u.FollowerCount.Add(1)); err != nil {
		return err
	}
	if _, err = udo.Where(u.ID.Eq(int32(userID))).Update(u.FollowCount, u.FollowCount.Add(1)); err != nil {
		return err
	}
	return tx.Commit()
}

func UnFollowUser(userID, toUserID int) (err error) {
	tq := query.Use(dal.DB)
	tx := tq.Begin() // 开启事务// 开启事务
	defer func() {
		if recover() != nil || err != nil {
			_ = tx.Rollback()
		}
	}()

	r := tx.Relation
	rdo := r.WithContext(context.Background())
	res, err := rdo.Where(r.UserID.Eq(int32(userID)), r.ToUserID.Eq(int32(toUserID))).Delete()
	if err != nil {
		return err
	}
	if res.RowsAffected > 0 {
		u := tx.User
		udo := u.WithContext(context.Background())
		if _, err = udo.Where(u.ID.Eq(int32(toUserID))).Update(u.FollowerCount, u.FollowerCount.Sub(1)); err != nil {
			return err
		}
		if _, err = udo.Where(u.ID.Eq(int32(userID))).Update(u.FollowCount, u.FollowCount.Sub(1)); err != nil {
			return err
		}
	}
	return tx.Commit()
}

func GetFollowList(userID int) ([]*UserInfo, error) {
	r := q.Relation
	rdo := r.WithContext(context.Background())
	resp, err := rdo.Where(r.UserID.Eq(int32(userID))).Preload(r.FollowUser).Find()
	if err != nil {
		return nil, err
	}

	list := make([]*UserInfo, len(resp))
	for i, v := range resp {
		list[i] = &UserInfo{
			ID:            int(v.ToUserID),
			Username:      v.FollowUser.Username,
			FollowCount:   int(v.FollowUser.FollowCount),
			FollowerCount: int(v.FollowUser.FollowerCount),
			IsFollow:      true,
		}
	}
	return list, nil
}
