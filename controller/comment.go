package controller

import (
	"github.com/chenshone/tiktok-lite/service"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

type commentActionReq struct {
	VideoID     int    `json:"video_id" form:"video_id" binding:"required"`
	CommentID   int    `json:"comment_id" form:"comment_id"`
	ActionType  int    `json:"action_type" form:"action_type" binding:"required"`
	CommentText string `json:"comment_text" form:"comment_text"`
}

type commentActionResp struct {
	Code    int         `json:"status_code"`
	Msg     string      `json:"status_msg"`
	Comment interface{} `json:"comment"`
}

func CommentAction(c *gin.Context) {
	userID := c.GetInt("user_id")
	req := &commentActionReq{}
	err := c.ShouldBind(req)
	log.Println("comment action req", req)
	if err != nil {
		c.JSON(200, &commentActionResp{
			Code: -1,
			Msg:  "参数不全/错误",
		})
		return
	}
	switch req.ActionType {
	case 1: // 添加评论
		data, err := service.PublishComment(userID, req.VideoID, req.CommentText)
		if err != nil {
			c.JSON(200, &commentActionResp{
				Code: -1,
				Msg:  "评论失败",
			})
			return
		}
		c.JSON(200, &commentActionResp{
			Code:    0,
			Msg:     "success",
			Comment: data,
		})
	case 2: // 删除评论
		err := service.RemoveComment(req.CommentID, req.VideoID)
		if err != nil {
			c.JSON(200, &commentActionResp{
				Code: -1,
				Msg:  "删除失败",
			})
			return
		}
		c.JSON(200, &commentActionResp{
			Code: 0,
			Msg:  "success",
		})
	default:
		c.JSON(200, &commentActionResp{
			Code: -1,
			Msg:  "参数错误",
		})
	}
}

type commentListResp struct {
	Code        int         `json:"status_code"`
	Msg         string      `json:"status_msg"`
	CommentList interface{} `json:"comment_list"`
}

func GetCommentList(c *gin.Context) {
	userID := c.GetInt("user_id")
	vid, ok := c.GetQuery("video_id")
	if !ok {
		c.JSON(200, &commentListResp{
			Code: -1,
			Msg:  "参数不匹配",
		})
		return
	}
	videoID, err := strconv.Atoi(vid)
	if err != nil {
		c.JSON(200, &commentListResp{
			Code: -1,
			Msg:  "参数错误",
		})
		return
	}
	list, err := service.GetCommentList(userID, videoID)
	if err != nil {
		c.JSON(200, &commentListResp{
			Code: -1,
			Msg:  "获取评论失败",
		})
		return
	}
	c.JSON(200, &commentListResp{
		Code:        0,
		Msg:         "success",
		CommentList: list,
	})
}
