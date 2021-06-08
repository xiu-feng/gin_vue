package middle

import (
	"github.com/gin-gonic/gin"
	"myself_rep/gin_vue/global"
	"myself_rep/gin_vue/model"
	"myself_rep/gin_vue/utils"
	"net/http"
	"strings"
)

//鉴权JWT中间件

// AuthMiddleware 认证
func AuthMiddleware()  gin.HandlerFunc{
	return func(ctx *gin.Context)  {
		// 获取 authorization hreader
		tokenString := ctx.GetHeader("Authorization")

		// 验证token格式
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer "){
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "message": "权限不足"})
			ctx.Abort()
			return
		}

		tokenString = tokenString[7:]
		token, claims, err := utils.ParseToken(tokenString)
		if err != nil || !token.Valid{
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "message": "权限不足"})
			ctx.Abort()
			return
		}

		// 验证通过后获取claim中的userID
		userID := claims.UserID
		DB := global.GVA_DB
		var user model.User
		DB.Table("t_user").First(&user, userID)

		// 用户不存在
		if user.ID == 0 {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "message": "权限不足"})
			ctx.Abort()
			return
		}

		// 用户存在,存入上下文
		ctx.Set("user", user)
		ctx.Next()
	}
}