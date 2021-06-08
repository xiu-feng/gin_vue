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