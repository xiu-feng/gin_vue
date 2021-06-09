package model

import "myself_rep/gin_vue/global"

type User struct {
	ID int `gorm:"column:id"`
	Name string `gorm:"column:name"`
	Password string `gorm:"column:password"`
}

type UserRepo interface {
	Add()
	Select() *[]User
	FindOne(where interface{}) *User
	Update(user *User)
	Delete(where interface{})
}

func (u *User) Add() {
	global.GVA_DB.Table("t_user").Create(&u)
}

func (u *User) Select() *[]User {
	users := []User{}
	global.GVA_DB.Table("t_user").Find(&users)
	return &users
}

func (u *User) FindOne(where interface{}) *User {
	global.GVA_DB.Table("t_user").Find(&u,where)
	return u
}
//u中携带id ｜ user为需要修改的对象数值
func (u *User) Update(user *User) {
	global.GVA_DB.Table("t_user").Model(&u).Update(&user)
}

//根据主键删除,不是软删除
func (u *User) Delete(where interface{})  {
	global.GVA_DB.Table("t_user").Delete(&u,where)
}