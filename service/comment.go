package service

import (
	"context"
	"errors"
	"github.com/chenshone/tiktok-lite/dal/model"
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

func PublishComment(userID, videoID int, content string) (*CommentInfo, error) {
	data := &model.Comment{
		UserID:   int32(userID),
		VideoID:  int32(videoID),
		Content:  content,
		CreateAt: time.Now(),
	}
	c := q.Comment
	cdo := c.WithContext(context.Background())
	err := cdo.Save(data)
	log.Println("save comment", data)
	if err != nil {
		return nil, err
	}

	u := q.User
	udo := u.WithContext(context.Background())
	users, err := udo.Where(u.ID.Eq(int32(userID))).Find()
	if err != nil || len(users) == 0 {
		return nil, errors.New("用户不存在")
	}

	return &CommentInfo{
		ID: int(data.ID),
		User: userinfo{
			ID:            int(users[0].ID),
			Name:          users[0].Username,
			FollowCount:   int(users[0].FollowerCount),
			FollowerCount: int(users[0].FollowerCount),
			IsFollow:      false,
		},
		Content:    data.Content,
		CreateDate: data.CreateAt.Format("01-02"),
	}, nil
}

func RemoveComment(commentID int) error {
	c := q.Comment
	cdo := c.WithContext(context.Background())
	_, err := cdo.Where(c.ID.Eq(int32(commentID))).Delete()
	return err
}

func GetCommentList(videoID int) ([]*CommentInfo, error) {
	c := q.Comment
	cdo := c.WithContext(context.Background())
	comments, err := cdo.Where(c.VideoID.Eq(int32(videoID))).Preload(c.Author).Find()
	if err != nil {
		return nil, err
	}
	commentList := make([]*CommentInfo, len(comments))
	for i, v := range comments {
		commentList[i] = &CommentInfo{
			ID: int(v.ID),
			User: userinfo{
				ID:            int(v.Author.ID),
				Name:          v.Author.Username,
				FollowCount:   int(v.Author.FollowCount),
				FollowerCount: int(v.Author.FollowerCount),
				IsFollow:      false,
			},
			Content:    v.Content,
			CreateDate: v.CreateAt.Format("01-02"),
		}
	}
	return commentList, nil
}
