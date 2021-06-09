package main

import (
	"github.com/gin-gonic/gin"
	"myself_rep/gin_vue/controller"
	"myself_rep/gin_vue/middle"
	"net/http"
)

func RouterList(r *gin.Engine) *gin.Engine{
	r.StaticFS("/static",http.Dir("static")) //配置静态资源
	r.StaticFile("favicon.ico","favicon.ico")
	r.LoadHTMLGlob("templates/*") //配置模版引擎
	//跨域配置
	r.Use(middle.Cors())
	//前端页面直接放行
	view := r.Group("/") //路由分组
	view.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK,"index.html",gin.H{"title":"前端模版首页"})
	})
	//后台路由放行
	api := r.Group("/api")
	api.POST("/login",controller.LoginHandle)
	//带jwt验证的路由
	authApi := api.Group("/", middle.AuthMiddleware())
	authApi.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK,gin.H{"msg":"首页"})
	})
	//用户路由
	authApi.GET("/user",controller.GetUser)
	authApi.POST("/user",controller.AddUser)
	authApi.GET("/user/:id",controller.FindOne)
	authApi.PUT("/user/:id",controller.UpdateUser)
	authApi.DELETE("/user/:id",controller.DeleteUser)

	return r
}
