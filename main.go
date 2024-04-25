package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// if err := dal.InitDB(); err != nil {
	// 	os.Exit(-1)
	// }
	// query.SetDefault(dal.DB)
	r := gin.Default()

	// router.InitRouter(r)

	// r.StaticFS("/assets", gin.Dir("./assets", true))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	err := r.Run(":8021")
	if err != nil {
		return
	}
}
