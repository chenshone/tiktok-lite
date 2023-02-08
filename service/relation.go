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
	// 放置重复取消关注
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

// GetFollowList 获取关注列表
func GetFollowList(userID, targetUserID int) ([]*UserInfo, error) {
	r := q.Relation
	rdo := r.WithContext(context.Background())
	resp, err := rdo.Where(r.UserID.Eq(int32(targetUserID))).Preload(r.FollowUser).Find()
	if err != nil {
		return nil, err
	}

	list := make([]*UserInfo, len(resp))
	isFollow := false
	for i, v := range resp {
		if _, err = rdo.Where(r.UserID.Eq(int32(userID)), r.ToUserID.Eq(v.ToUserID)).First(); err == nil {
			isFollow = true
		} else {
			isFollow = false
		}
		list[i] = &UserInfo{
			ID:            int(v.ToUserID),
			Username:      v.FollowUser.Username,
			FollowCount:   int(v.FollowUser.FollowCount),
			FollowerCount: int(v.FollowUser.FollowerCount),
			IsFollow:      isFollow,
		}
	}
	return list, nil
}

// GetFollowerList 获取粉丝列表
func GetFollowerList(userID, targetUserID int) ([]*UserInfo, error) {
	r := q.Relation
	rdo := r.WithContext(context.Background())
	resp, err := rdo.Where(r.ToUserID.Eq(int32(targetUserID))).Preload(r.User).Find()
	if err != nil {
		return nil, err
	}

	list := make([]*UserInfo, len(resp))
	isFollow := false
	for i, v := range resp {
		if _, err = rdo.Where(r.UserID.Eq(int32(userID)), r.ToUserID.Eq(v.UserID)).First(); err == nil {
			isFollow = true
		} else {
			isFollow = false
		}
		list[i] = &UserInfo{
			ID:            int(v.UserID),
			Username:      v.User.Username,
			FollowCount:   int(v.User.FollowCount),
			FollowerCount: int(v.User.FollowerCount),
			IsFollow:      isFollow,
		}
	}
	return list, nil
}

// GetFriendList 获取好友列表，好友列表是指双方互相关注的用户
func GetFriendList(userID int) ([]*UserInfo, error) {
	r := q.Relation
	rdo := r.WithContext(context.Background())
	// 获取关注列表
	followIDs, err := rdo.Select(r.ToUserID).Where(r.UserID.Eq(int32(userID))).Find()
	if err != nil {
		return nil, err
	}
	followIDsSlice := make([]int32, len(followIDs))
	for i, v := range followIDs {
		followIDsSlice[i] = v.ToUserID
	}
	// 筛选出双方互相关注的用户
	resp, err := rdo.Where(r.UserID.In(followIDsSlice...), r.ToUserID.Eq(int32(userID))).Preload(r.User).Find()
	if err != nil {
		return nil, err
	}
	list := make([]*UserInfo, len(resp))
	for i, v := range resp {
		list[i] = &UserInfo{
			ID:            int(v.UserID),
			Username:      v.User.Username,
			FollowCount:   int(v.User.FollowCount),
			FollowerCount: int(v.User.FollowerCount),
			IsFollow:      true,
		}
	}
	return list, nil
}
