package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CollectRoute(router *gin.Engine) *gin.Engine {
	router.StaticFS("static", http.Dir("./static"))
	router.StaticFile("/favicon.ico", "./favicon.ico")
	//配置模版 配置路由分组
	router.LoadHTMLGlob("templates/*")
	view := router.Group("/")
	api := router.Group("api/")
	//后台页面路由
	api.GET("/regiest", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"msg": "注册",
		})
	})

	//前端页面跳转

	view.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{"title": "我是前端的首页"})
	})

	return router
}
