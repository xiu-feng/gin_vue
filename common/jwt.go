package common

import (
	"gin_vue/model"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var SigningKey = []byte("xiufengifbduowibcjas")

type CustomClaims struct {
	UserID uint
	Name   string
	jwt.StandardClaims
}

/*
jwt生成身份令牌，需传入结构体
*/
func GenerateToken(user *model.User) string {
	nowTime := time.Now()
	expireTime := nowTime.Add(3 * time.Hour)
	// Create the Claims
	claims := CustomClaims{
		user.ID,
		user.Name,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "test",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString(SigningKey)
	return tokenString
}

/*
jwt 解析身份令啊，传入token字符串
*/
func ParseToken(tokenString string) (*CustomClaims, error) {
	tokenClaims, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return SigningKey, nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*CustomClaims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
