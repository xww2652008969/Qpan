package utils

import (
	"Qpan/global"
	"github.com/dgrijalva/jwt-go"
	"time"
)

//创建jwt
type Myclains struct {
	Name string `json:"name"`
	jwt.StandardClaims
}

// Createtoken 传入用户名 返回加密后的token 和error
func Createtoken(name string) (string, error) {
	var c = Myclains{
		name,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Unix() + 100, //结束时间
			Issuer:    "xww",
			NotBefore: time.Now().Unix(), //开始时间
		},
	}
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	token, err := t.SignedString(global.JwtKey)
	if err != nil {
		return "", err
	} else {
		return token, err
	}
}

// Parsetoken 传入token 检验是否通过
func Parsetoken(token string) (string, error) {
	t, err := jwt.ParseWithClaims(token, &Myclains{}, func(token *jwt.Token) (interface{}, error) {
		return global.JwtKey, nil
	})
	if err != nil {
		return "", err
	} else {
		d := t.Claims.(*Myclains).Name
		return d, nil
	}
}
