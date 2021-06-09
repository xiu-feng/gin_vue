package controller

import (
	"github.com/gin-gonic/gin"
	"myself_rep/gin_vue/global"
	"myself_rep/gin_vue/model"
	"net/http"
	"strconv"
)

func GetUser(c *gin.Context) {
    user :=	model.User{}
    users:= user.Select()
	c.JSON(200,gin.H{"data":users})
}

func AddUser(c *gin.Context){
	j := make(map[string]interface{})
	c.ShouldBindJSON(&j)
	name :=j["name"]
	password := j["password"]
	user := &model.User{
		Name: name.(string),
		Password: password.(string),
	}
	user.Add()
	c.JSON(http.StatusOK,gin.H{"code":global.SUCCESS,"msg":"添加成功"})
}

func FindOne (c *gin.Context){
	id := c.Param("id")
	u  := &model.User{}
	user := u.FindOne(id)
	c.JSON(http.StatusOK,gin.H{"code":global.SUCCESS,"msg":"查询成功","data":user})
}

func UpdateUser (c *gin.Context){
	ids := c.Param("id")
	id,_ := strconv.Atoi(ids)
	u :=&model.User{ID: id}

	j := make(map[string]interface{})
	c.ShouldBindJSON(&j)
	name :=j["name"]
	password := j["password"]
	user := model.User{Name: name.(string),Password: password.(string)}
	u.Update(&user)
	c.JSON(http.StatusOK,gin.H{"code":global.SUCCESS,"msg":"修改成功"})

}

func DeleteUser (c *gin.Context){
	id := c.Param("id")
	u :=&model.User{}
	u.Delete(id)
	c.JSON(http.StatusOK,gin.H{"code":global.SUCCESS,"msg":"删除成功"})
}