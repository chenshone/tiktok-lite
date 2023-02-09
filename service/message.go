package service

import (
	"context"
	"github.com/chenshone/tiktok-lite/dal/model"
)

func SendMessage(from, to int, content string) (err error) {
	m := q.Message
	mdo := m.WithContext(context.Background())
	data := &model.Message{
		UserID:   int32(from),
		ToUserID: int32(to),
		Content:  content,
	}
	err = mdo.Create(data)
	if err != nil {
		return err
	}
	return nil
}

type message struct {
	Id         int    `json:"id"`
	Content    string `json:"content"`
	CreateTime string `json:"create_time"`
}

func GetMessageList(from, to int) ([]*message, error) {
	m := q.Message
	mdo := m.WithContext(context.Background())
	resp, err := mdo.Where(m.UserID.Eq(int32(from)), m.ToUserID.Eq(int32(to))).Order(m.CreateAt).Find()
	if err != nil {
		return nil, err
	}
	list := make([]*message, len(resp))
	for i, v := range resp {
		list[i] = &message{
			Id:         int(v.ID),
			Content:    v.Content,
			CreateTime: v.CreateAt.Format("2006-01-02 15:04:05"),
		}
	}
	return list, nil
}
