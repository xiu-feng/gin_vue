package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RouterList(r *gin.Engine) *gin.Engine{
	r.StaticFS("/static",http.Dir("static"))
	r.StaticFile("favicon.ico","favicon.ico")
	r.LoadHTMLGlob("templates/*")
	view := r.Group("/")
	api := r.Group("/api")
	view.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK,"index.html",gin.H{"title":"前端模版首页"})
	})
	api.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK,gin.H{"msg":"api首页"})
	})
	return r
}
