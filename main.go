package main

import (
	"gin_vue/controller"

	"github.com/gin-gonic/gin"
)

// 启动入口

func main() {
	r := gin.Default()
	r.GET("/ping", controller.Ping)
	r = CollectRoute(r)
	r.Run(":9000")
}
