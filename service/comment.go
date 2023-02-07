package service

import (
	"context"
	"errors"
	"github.com/chenshone/tiktok-lite/dal"
	"github.com/chenshone/tiktok-lite/dal/model"
	"github.com/chenshone/tiktok-lite/dal/query"
	"log"
	"time"
)

type CommentInfo struct {
	ID         int      `json:"id"`
	User       userinfo `json:"user"`
	Content    string   `json:"content"`
	CreateDate string   `json:"create_date"`
}

type userinfo struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	FollowCount   int    `json:"follow_count"`
	FollowerCount int    `json:"follower_count"`
	IsFollow      bool   `json:"is_follow"`
}

func PublishComment(userID, videoID int, content string) (res *CommentInfo, err error) {
	tq := query.Use(dal.DB)
	tx := tq.Begin() // 开启事务
	defer func() {
		if recover() != nil || err != nil {
			_ = tx.Rollback()
		}
	}()
	data := &model.Comment{
		UserID:   int32(userID),
		VideoID:  int32(videoID),
		Content:  content,
		CreateAt: time.Now(),
	}
	c := tx.Comment
	cdo := c.WithContext(context.Background())
	log.Println("save comment", data)
	err = cdo.Save(data)
	if err != nil {
		return nil, err
	}

	u := tx.User
	udo := u.WithContext(context.Background())
	users, err := udo.Where(u.ID.Eq(int32(userID))).Find()
	if err != nil || len(users) == 0 {
		return nil, errors.New("用户不存在")
	}
	v := tx.Video
	vdo := v.WithContext(context.Background())
	_, err = vdo.Where(v.ID.Eq(int32(videoID))).Update(v.CommentCount, v.CommentCount.Add(1))
	if err != nil {
		return nil, err
	}
	return &CommentInfo{
		ID: int(data.ID),
		User: userinfo{
			ID:            int(users[0].ID),
			Name:          users[0].Username,
			FollowCount:   int(users[0].FollowerCount),
			FollowerCount: int(users[0].FollowerCount),
			IsFollow:      true,
		},
		Content:    data.Content,
		CreateDate: data.CreateAt.Format("01-02"),
	}, tx.Commit()
}

func RemoveComment(commentID, videoID int) (err error) {
	tq := query.Use(dal.DB)
	tx := tq.Begin() // 开启事务
	defer func() {
		if recover() != nil || err != nil {
			_ = tx.Rollback()
		}
	}()
	c := tx.Comment
	cdo := c.WithContext(context.Background())
	if _, err = cdo.Where(c.ID.Eq(int32(commentID))).Delete(); err != nil {
		return err
	}
	v := tx.Video
	vdo := v.WithContext(context.Background())
	if _, err = vdo.Where(v.ID.Eq(int32(videoID))).Update(v.CommentCount, v.CommentCount.Sub(1)); err != nil {
		return err
	}
	return tx.Commit()
}

func GetCommentList(userID, videoID int) ([]*CommentInfo, error) {
	c := q.Comment
	cdo := c.WithContext(context.Background())
	comments, err := cdo.Where(c.VideoID.Eq(int32(videoID))).Preload(c.Author).Find()
	if err != nil {
		return nil, err
	}
	commentList := make([]*CommentInfo, len(comments))
	r := q.Relation
	rdo := r.WithContext(context.Background())
	var isFollow bool
	for i, v := range comments {
		if _, err := rdo.Where(r.UserID.Eq(int32(userID)), r.ToUserID.Eq(v.Author.ID)).First(); err == nil {
			isFollow = true
		} else {
			isFollow = false
		}
		commentList[i] = &CommentInfo{
			ID: int(v.ID),
			User: userinfo{
				ID:            int(v.Author.ID),
				Name:          v.Author.Username,
				FollowCount:   int(v.Author.FollowCount),
				FollowerCount: int(v.Author.FollowerCount),
				IsFollow:      isFollow,
			},
			Content:    v.Content,
			CreateDate: v.CreateAt.Format("01-02"),
		}
	}
	return commentList, nil
}
