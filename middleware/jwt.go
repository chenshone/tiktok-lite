package middleware

import (
	"github.com/chenshone/tiktok-lite/controller"
	"github.com/chenshone/tiktok-lite/util/util"
	"github.com/gin-gonic/gin"
)

func JWTAuth() func(c *gin.Context) {
	return func(c *gin.Context) {
		userId := c.Query("user_id")
		token := c.Query("token")
		if userId == "" || token == "" {
			c.JSON(200, &controller.UserInfoRes{
				Code: -1,
				Msg:  "缺失参数",
				User: nil,
			})
			c.Abort()
			return
		}
		jwt := util.JWT{}
		claims, err := jwt.ParseToken(token)
		if err != nil || claims != userId {
			c.JSON(200, &controller.UserInfoRes{
				Code: -1,
				Msg:  "token验证失败",
				User: nil,
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
