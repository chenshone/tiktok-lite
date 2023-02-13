package service

import (
	"context"
	"github.com/chenshone/tiktok-lite/dal/model"
	"sort"
	"strconv"
)

// 因为客户端采用轮询的方式获取消息，所以需要一个map来记录当前聊天的两个用户是否发送新消息
// 判断当前聊天的两个用户是否发送新消息
var messageIsUpdate = make(map[string]bool)

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
	messageIsUpdate[strconv.Itoa(from)+"-"+strconv.Itoa(to)] = true
	return nil
}

type message struct {
	Id         int    `json:"id"`
	ToUserId   int    `json:"to_user_id"`
	FromUserId int    `json:"from_user_id"`
	Content    string `json:"content"`
	CreateTime int64  `json:"create_time"`
}

func GetMessageList(from, to int) ([]*message, error) {
	if st, ok := messageIsUpdate[strconv.Itoa(from)+"-"+strconv.Itoa(to)]; ok && !st { // 存在且为false
		return nil, nil
	}
	m := q.Message
	mdo := m.WithContext(context.Background())
	resp1, err := mdo.Where(m.UserID.Eq(int32(from)), m.ToUserID.Eq(int32(to))).Find()
	resp2, err := mdo.Where(m.UserID.Eq(int32(to)), m.ToUserID.Eq(int32(from))).Find()
	if err != nil {
		return nil, err
	}
	list := make([]*message, len(resp1)+len(resp2))
	for i, v := range resp1 {
		list[i] = &message{
			Id:         int(v.ID),
			ToUserId:   int(v.ToUserID),
			FromUserId: int(v.UserID),
			Content:    v.Content,
			CreateTime: v.CreateAt.Unix(),
		}
	}
	for i, v := range resp2 {
		list[i+len(resp1)] = &message{
			Id:         int(v.ID),
			ToUserId:   int(v.ToUserID),
			FromUserId: int(v.UserID),
			Content:    v.Content,
			CreateTime: v.CreateAt.Unix(),
		}
	}
	sort.Slice(list, func(i, j int) bool {
		return list[i].CreateTime < list[j].CreateTime
	})
	messageIsUpdate[strconv.Itoa(from)+"-"+strconv.Itoa(to)] = false
	return list, nil
}
