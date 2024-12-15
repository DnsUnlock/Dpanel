package jwt

import "github.com/golang-jwt/jwt"

type Claims struct {
	Token string
	jwt.StandardClaims
}

type ClaimsInfo struct {
	Token
	jwt.StandardClaims
}

type Token struct {
	Uuid      string // 用户唯一标识
	UserAgent string // 客户端请求头
	Finger    string // 令牌指纹
}
