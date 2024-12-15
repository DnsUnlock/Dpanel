package jwe

import (
	"encoding/json"
	"errors"
	jwt2 "github.com/DnsUnlock/Dpanel/backend/model/jwt"
	"github.com/DnsUnlock/Dpanel/backend/utils/aes"
	"github.com/golang-jwt/jwt"
	"time"
)

var jweKey = []byte("0ee2662765b978639f503d90e95f8fa7")

func SetToken(userInfo jwt2.Token) (string, error) {
	marshal, err := json.Marshal(userInfo)
	if err != nil {
		return "", err
	}
	userInfoAes, err := aes.EncryptAES(jweKey, string(marshal))
	if err != nil {
		return "", err
	}
	claims := &jwt2.Claims{
		Token: userInfoAes,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), // 强制过期时间
			IssuedAt:  time.Now().Unix(),
			Issuer:    "DnsUnlock.com",  // 签名颁发者
			Subject:   "User Web Token", // 签名主题
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jweKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func GetToken(tokenString string) (c jwt2.ClaimsInfo, err error) {
	var d jwt2.Claims
	token, err := jwt.ParseWithClaims(tokenString, &d, func(token *jwt.Token) (interface{}, error) {
		return jweKey, nil
	})
	if err != nil {
		return c, err
	}
	// 判断token是否有效
	if token.Valid {
		//解析token字段
		userInfo, err := aes.DecryptAES(jweKey, d.Token)
		if err != nil {
			return c, err
		}
		var t jwt2.Token
		err = json.Unmarshal([]byte(userInfo), &t)
		if err != nil {
			return c, err
		}
		c.Token = t
		return c, nil
	} else {
		return c, errors.New("token 无效")
	}
}
