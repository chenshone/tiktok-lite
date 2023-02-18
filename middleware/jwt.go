package middleware

import (
	"context"
	"github.com/chenshone/tiktok-lite/dal/query"
	"github.com/chenshone/tiktok-lite/util/util"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

type JWT struct {
	Token string `form:"token" json:"token"`
}

func JWTAuth() func(c *gin.Context) {
	return func(c *gin.Context) {
		jwtToken := JWT{}
		log.Println("URL: ", c.Request.URL.Path)
		isSkip := false
		// 因为feed视频流接口也需要token获取当前用户的id，但是feed接口不需要token验证
		if c.Request.URL.Path == "/douyin/feed/" {
			isSkip = true
		}
		err := c.Bind(&jwtToken)
		if err != nil {
			if isSkip {
				c.Next()
				return
			}
			c.JSON(200, &gin.H{
				"status_code": -1,
				"status_msg":  "缺失参数",
			})
			c.Abort()
			return
		}
		token := jwtToken.Token
		if token == "" {
			if isSkip {
				c.Next()
				return
			}
			c.JSON(200, &gin.H{
				"status_code": -1,
				"status_msg":  "缺失参数",
			})
			c.Abort()
			return
		}
		jwt := util.JWT{}
		claim, err := jwt.ParseToken(token)
		if err != nil {
			if isSkip {
				c.Next()
				return
			}
			c.JSON(200, &gin.H{
				"status_code": -1,
				"status_msg":  "token验证失败",
			})
			c.Abort()
			return
		}
		// 获取用户id
		userId, err := strconv.Atoi(claim)
		if err != nil {
			if isSkip {
				c.Next()
				return
			}
			c.JSON(200, &gin.H{
				"status_code": -1,
				"status_msg":  "token验证失败",
			})
			c.Abort()
			return
		}
		// 判断用户是否存在
		u := query.Q.User
		udo := u.WithContext(context.Background())
		find, err := udo.Where(u.ID.Eq(int32(userId))).Find()
		if err != nil {
			if isSkip {
				c.Next()
				return
			}
			c.JSON(200, &gin.H{
				"status_code": -1,
				"status_msg":  "token验证失败",
			})
			c.Abort()
			return
		}
		if len(find) == 0 {
			if isSkip {
				c.Next()
				return
			}
			c.JSON(200, &gin.H{
				"status_code": -1,
				"status_msg":  "token验证失败",
			})
			c.Abort()
			return
		}
		// 将用户id存入上下文
		c.Set("user_id", userId)
		c.Next()
	}
}
