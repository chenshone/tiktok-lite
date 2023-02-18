package service

import (
	"context"
	"github.com/chenshone/tiktok-lite/dal"
	"github.com/chenshone/tiktok-lite/dal/model"
	"github.com/chenshone/tiktok-lite/dal/query"
	"log"
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
			ID:              int(v.ToUserID),
			Username:        v.FollowUser.Username,
			FollowCount:     int(v.FollowUser.FollowCount),
			FollowerCount:   int(v.FollowUser.FollowerCount),
			IsFollow:        isFollow,
			Avatar:          v.FollowUser.Avatar,
			BackgroundImage: v.FollowUser.BackgroundImage,
			Signature:       v.FollowUser.Signature,
			TotalFavorited:  int(v.FollowUser.TotalFavorited),
			WorkCount:       int(v.FollowUser.WorkCount),
			FavoriteCount:   int(v.FollowUser.FavoriteCount),
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
			ID:              int(v.UserID),
			Username:        v.User.Username,
			FollowCount:     int(v.User.FollowCount),
			FollowerCount:   int(v.User.FollowerCount),
			IsFollow:        isFollow,
			Avatar:          v.User.Avatar,
			BackgroundImage: v.User.BackgroundImage,
			Signature:       v.User.Signature,
			TotalFavorited:  int(v.User.TotalFavorited),
			WorkCount:       int(v.User.WorkCount),
			FavoriteCount:   int(v.User.FavoriteCount),
		}
	}
	return list, nil
}

type FriendInfo struct {
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
	Message         string `json:"message"` // 最新消息
	MsgType         int    `json:"msgType"` // 消息类型 0 =>当前清求用户接收的消息，1=>当前请求用户发送的消息
}

// GetFriendList 获取好友列表，好友列表是指双方互相关注的用户
func GetFriendList(userID int) ([]*FriendInfo, error) {
	log.Println("GetFriendList ", userID)
	r := q.Relation
	rdo := r.WithContext(context.Background())
	m := q.Message
	mdo := m.WithContext(context.Background())
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
	list := make([]*FriendInfo, len(resp))
	for i, v := range resp {
		// 获取最新消息
		msgs, err := mdo.Where(m.UserID.Eq(int32(userID)), m.ToUserID.Eq(v.UserID)).Or(m.UserID.Eq(v.UserID),
			m.ToUserID.Eq(int32(userID))).Order(m.CreateAt.Desc()).Find()
		if err != nil {
			return nil, err
		}

		var msg *model.Message = new(model.Message)
		if len(msgs) > 0 {
			msg = msgs[0]
		}
		msgType := 0
		//1=>当前请求用户发送的消息
		if msg.UserID == int32(userID) {
			msgType = 1
		}
		list[i] = &FriendInfo{
			ID:              int(v.UserID),
			Username:        v.User.Username,
			FollowCount:     int(v.User.FollowCount),
			FollowerCount:   int(v.User.FollowerCount),
			IsFollow:        true,
			Avatar:          v.User.Avatar,
			BackgroundImage: v.User.BackgroundImage,
			Signature:       v.User.Signature,
			TotalFavorited:  int(v.User.TotalFavorited),
			WorkCount:       int(v.User.WorkCount),
			FavoriteCount:   int(v.User.FavoriteCount),
			Message:         msg.Content,
			MsgType:         msgType,
		}
	}
	log.Printf("user %d friend list: %v", userID, list)
	return list, nil
}
