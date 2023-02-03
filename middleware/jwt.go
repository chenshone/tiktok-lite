package middleware

import (
	"github.com/chenshone/tiktok-lite/util/util"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

type JWT struct {
	Token string `form:"token" json:"token" binding:"required"`
}

func JWTAuth() func(c *gin.Context) {
	return func(c *gin.Context) {
		jwtToken := JWT{}
		err := c.Bind(&jwtToken)
		if err != nil {
			c.JSON(200, &gin.H{
				"status_code": -1,
				"status_msg":  "缺失参数",
			})
			c.Abort()
			return
		}
		token := jwtToken.Token
		if token == "" {
			c.JSON(200, &gin.H{
				"status_code": -1,
				"status_msg":  "缺失参数",
			})
			c.Abort()
			return
		}
		jwt := util.JWT{}
		claim, err1 := jwt.ParseToken(token)
		userId, err2 := strconv.Atoi(claim)
		log.Println(claim)
		log.Println(token)
		if err1 != nil || err2 != nil {
			c.JSON(200, &gin.H{
				"status_code": -1,
				"status_msg":  "token验证失败",
			})
			c.Abort()
			return
		}
		c.Set("user_id", userId)
		c.Next()
	}
}
