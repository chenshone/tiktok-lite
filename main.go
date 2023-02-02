package main

import (
	"github.com/chenshone/tiktok-lite/dal"
	"github.com/chenshone/tiktok-lite/dal/query"
	"github.com/chenshone/tiktok-lite/router"
	"github.com/gin-gonic/gin"
	"os"
)

func main() {
	if err := dal.InitDB(); err != nil {
		os.Exit(-1)
	}
	query.SetDefault(dal.DB)
	r := gin.Default()
	r.Use(gin.Logger())

	router.InitRouter(r)

	err := r.Run()
	if err != nil {
		return
	}
}
