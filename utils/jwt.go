package utils

import (
	"github.com/dgrijalva/jwt-go"
	"myself_rep/gin_vue/global"
	"myself_rep/gin_vue/model"
	"time"
)

//生成jwt，验证jwt

// 服务端加密字符串
var jwtKey = []byte(global.GVA_CONFIG.Jwt.SignKey)

// Claims jwt，UserID为自定义字段，jwt.StandardClaims包含官方字段
type Claims struct {
	UserID int
	jwt.StandardClaims
}

// ReleaseToken 发放token, model.User为自定义模型
func ReleaseToken(user *model.User) (string, error) {
	expiresTime := time.Now().Add(7 * 24 * time.Hour)
	claims := &Claims{
		UserID: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiresTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "gingo",
			Subject:   "user token",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// ParseToken 解析token
func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (i interface{}, err error) {
		return jwtKey, nil
	})

	return token, claims, err
}