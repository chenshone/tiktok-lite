package controller

import (
	"github.com/chenshone/tiktok-lite/service"
	"github.com/gin-gonic/gin"
	"strconv"
)

func FollowUserOrNot(c *gin.Context) {
	uid := c.GetInt("user_id")
	followID, ok1 := c.GetQuery("to_user_id")
	action_type, ok2 := c.GetQuery("action_type")
	if !ok1 || !ok2 {
		c.JSON(200, &gin.H{
			"status_code": -1,
			"status_msg":  "参数不全",
		})
		return
	}
	fid, err1 := strconv.Atoi(followID)
	action, err2 := strconv.Atoi(action_type)
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
