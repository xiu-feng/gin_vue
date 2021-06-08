package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"myself_rep/gin_vue/global"
	"myself_rep/gin_vue/model"
	"myself_rep/gin_vue/utils"
	"net/http"
)

func LoginHandle(c *gin.Context){
	j := make(map[string]interface{})
	c.ShouldBindJSON(&j)
	//验证传入字符串是否合法(是否为空，格式前端验证)
	//前端确定传入参数非空、符合6-12位、符合字符串与数字符号的两种组合
	name := j["name"]
	password := j["password"]
	if name == "" || password =="" {
		c.JSON(http.StatusOK,gin.H{"msg":"用户名或密码不能为空！","code":global.LOGINNULL})
		return
	} else if len([]rune(name.(string))) < 6 || len([]rune(name.(string))) > 12 ||
		len([]rune(password.(string))) < 6 || len([]rune(password.(string))) > 12 {
		c.JSON(http.StatusOK,gin.H{"msg":"用户名或密码不合法","code":global.LOGINLEN})
		return
	} else {
		//查找数据库，验证用户名密码
		user := model.User{}
		res := user.FindOne(j)
		if res.ID == 0 {
			c.JSON(http.StatusOK,gin.H{"msg":"用户或密码错误","code":global.LOGINUS})
		} else {
			//生成token，返回
			token,err :=utils.ReleaseToken(res)
			if err!= nil {
				fmt.Printf("ReleaseToken faild %s\n",err)
			}
			c.JSON(http.StatusOK,gin.H{"msg":"登陆成功","code":global.SUCCESS,"token":token})
		}
	}
}
