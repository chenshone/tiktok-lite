package controller

import (
	"github.com/chenshone/tiktok-lite/service"
	"github.com/gin-gonic/gin"
	"strconv"
)

func FollowUserOrNot(c *gin.Context) {
	uid := c.GetInt("user_id")
	followID, ok1 := c.GetQuery("to_user_id")
	actionType, ok2 := c.GetQuery("action_type")
	if !ok1 || !ok2 {
		c.JSON(200, &gin.H{
			"status_code": -1,
			"status_msg":  "参数不全",
		})
		return
	}
	fid, err1 := strconv.Atoi(followID)
	action, err2 := strconv.Atoi(actionType)
	if err1 != nil || err2 != nil {
		c.JSON(200, &gin.H{
			"status_code": -1,
			"status_msg":  "参数错误",
		})
		return
	}
	switch action {
	case 1: // follow
		err := service.FollowUser(uid, fid)
		if err != nil {
			c.JSON(200, &gin.H{
				"status_code": -1,
				"status_msg":  err.Error(),
			})
			return
		}
		c.JSON(200, &gin.H{
			"status_code": 0,
			"status_msg":  "success",
		})
	case 2: // unfollow
		err := service.UnFollowUser(uid, fid)
		if err != nil {
			c.JSON(200, &gin.H{
				"status_code": -1,
				"status_msg":  err.Error(),
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

func GetFollowList(c *gin.Context) {
	uid, ok := c.GetQuery("user_id")
	if !ok {
		c.JSON(200, &gin.H{
			"status_code": -1,
			"status_msg":  "参数不全",
		})
		return
	}
	targetUserID, err := strconv.Atoi(uid)
	if err != nil {
		c.JSON(200, &gin.H{
			"status_code": -1,
			"status_msg":  "参数错误",
		})
		return
	}
	id := c.GetInt("user_id")
	followList, err := service.GetFollowList(id, targetUserID)
	if err != nil {
		c.JSON(200, &gin.H{
			"status_code": -1,
			"status_msg":  err.Error(),
		})
		return
	}
	c.JSON(200, &gin.H{
		"status_code": 0,
		"status_msg":  "success",
		"user_list":   followList,
	})
}

// GetFollowerList 获取粉丝列表
func GetFollowerList(c *gin.Context) {
	uid, ok := c.GetQuery("user_id")
	if !ok {
		c.JSON(200, &gin.H{
			"status_code": -1,
			"status_msg":  "参数不全",
		})
		return
	}
	targetUserID, err := strconv.Atoi(uid)
	if err != nil {
		c.JSON(200, &gin.H{
			"status_code": -1,
			"status_msg":  "参数错误",
		})
		return
	}
	id := c.GetInt("user_id")
	followerList, err := service.GetFollowerList(id, targetUserID)
	if err != nil {
		c.JSON(200, &gin.H{
			"status_code": -1,
			"status_msg":  err.Error(),
		})
		return
	}
	c.JSON(200, &gin.H{
		"status_code": 0,
		"status_msg":  "success",
		"user_list":   followerList,
	})
}

func GetFriendList(c *gin.Context) {
	uid, ok := c.GetQuery("user_id")
	if !ok {
		c.JSON(200, &gin.H{
			"status_code": -1,
			"status_msg":  "参数不全",
		})
		return
	}
	targetUserID, err := strconv.Atoi(uid)
	if err != nil {
		c.JSON(200, &gin.H{
			"status_code": -1,
			"status_msg":  "参数错误",
		})
		return
	}
	id := c.GetInt("user_id")
	if id != targetUserID {
		c.JSON(200, &gin.H{
			"status_code": -1,
			"status_msg":  "不允许获取其他用户的好友列表",
		})
		return
	}
	friendList, err := service.GetFriendList(targetUserID)
	if err != nil {
		c.JSON(200, &gin.H{
			"status_code": -1,
			"status_msg":  err.Error(),
		})
		return
	}
	c.JSON(200, &gin.H{
		"status_code": 0,
		"status_msg":  "success",
		"user_list":   friendList,
	})
}
