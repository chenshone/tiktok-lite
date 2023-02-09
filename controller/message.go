package controller

import (
	"github.com/chenshone/tiktok-lite/service"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
	"strings"
)

type sendMessageReq struct {
	ToUserId   int    `json:"to_user_id" form:"to_user_id"`
	Content    string `json:"content" form:"content"`
	ActionType int    `json:"action_type" form:"action_type"`
}

func SendMessage(c *gin.Context) {
	uid := c.GetInt("user_id")
	var req sendMessageReq
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(200, &gin.H{
			"status_code": -1,
			"status_msg":  "参数错误",
		})
		return
	}
	req.Content = strings.Trim(req.Content, " ")
	if req.Content == "" {
		c.JSON(200, &gin.H{
			"status_code": -1,
			"status_msg":  "消息不能为空",
		})
		return
	}
	if uid == req.ToUserId {
		c.JSON(200, &gin.H{
			"status_code": -1,
			"status_msg":  "不能给自己发消息",
		})
		return
	}
	switch req.ActionType {
	case 1:
		err := service.SendMessage(uid, req.ToUserId, req.Content)
		if err != nil {
			c.JSON(200, &gin.H{
				"status_code": -1,
				"status_msg":  "发送失败",
			})
			return
		}
		c.JSON(200, &gin.H{
			"status_code": 0,
			"status_msg":  "success",
		})
	default:
		c.JSON(200, &gin.H{
			"status_code": -1,
			"status_msg":  "参数错误",
		})
	}

}

func GetMessageList(c *gin.Context) {
	uid := c.GetInt("user_id")
	toUserIdStr, ok := c.GetQuery("to_user_id")
	if !ok {
		c.JSON(200, &gin.H{
			"status_code": -1,
			"status_msg":  "缺少参数",
		})
		return
	}
	toUserId, err := strconv.Atoi(toUserIdStr)
	if err != nil {
		c.JSON(200, &gin.H{
			"status_code": -1,
			"status_msg":  "参数错误",
		})
		return
	}
	// 不能获取自己发给自己的消息
	if uid == toUserId {
		c.JSON(200, &gin.H{
			"status_code": -1,
			"status_msg":  "参数错误",
		})
		return
	}
	list, err := service.GetMessageList(uid, toUserId)
	if err != nil {
		c.JSON(200, &gin.H{
			"status_code": -1,
			"status_msg":  "获取失败",
		})
		return
	}
	log.Printf("获取消息列表成功: %#v", list[0])
	c.JSON(200, &gin.H{
		"status_code":  0,
		"status_msg":   "success",
		"message_list": list,
	})
}
