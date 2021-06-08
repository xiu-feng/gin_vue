package main

import (
	"github.com/gin-gonic/gin"
	_ "myself_rep/gin_vue/core"
	"myself_rep/gin_vue/global"
)


func main() {
	r := gin.Default()

	r = RouterList(r)
	r.Run(":"+global.GVA_CONFIG.Sys.HttpPort)
}