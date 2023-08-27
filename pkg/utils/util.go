package utils

import (
	"time"

	"chitchat2.0/pkg/setting"
	"github.com/dgrijalva/jwt-go"
)

var JWTsecret = []byte(setting.JWT_SECRET)

// Claims 是 Token 的要求
type Claims struct {
	Id       uint   `json:"id"`
	UserName string `json:"user_name"`
	Password string `json:"password"`
	jwt.StandardClaims
}

// GenerateToken 签发 Token
func GenerateToken(id uint, username, password string) (string, error) {
	// 获取当前时间
	nowTime := time.Now()
	// 登录24小时之后的时间
	expireTime := nowTime.Add(24 * time.Hour)

	// token 的要求
	claims := Claims{
		Id:       id,
		UserName: username,
		Password: password,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "chitchat2.0",
		},
	}
	tokenCliams := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenCliams.SignedString(JWTsecret)
	return token, err
}
