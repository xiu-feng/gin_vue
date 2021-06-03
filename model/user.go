package model

import (
	"gin_vue/common"

	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name     string
	Password string
}

func (u *User) Create() {
	DB := common.GetDB()
	DB.AutoMigrate(&u)
	DB.Create(&u)
}

func (u *User) SelectAll() interface{} {
	DB := common.GetDB()
	users := []User{}
	DB.Find(&users)
	return users
}

func (u *User) Update(id *int) {
	DB := common.GetDB()
	DB.Find(&u, &id)
	DB.Model(&u).Updates(&u)
}

func (u *User) Delete(id *int) {
	DB := common.GetDB()
	DB.Find(&u, &id)
	DB.Delete(&u)
}

func (u *User) Find(where interface{}) interface{} {
	DB := common.GetDB()
	DB.Find(&u, where)
	return &u
}

//TODO 模糊查询接口具体实现，分页查询具体实现
func (u *User) Like(where interface{}) interface{} {
	return nil
}

func (u *User) Page(where interface{}) interface{} {
	return nil
}
