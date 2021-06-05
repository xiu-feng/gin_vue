package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK,gin.H{"msg":"pong"})
	})
	r = RouterList(r)
	r.Run(":9000")
}