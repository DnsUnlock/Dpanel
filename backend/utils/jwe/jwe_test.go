package jwe

import (
	"github.com/DnsUnlock/Dpanel/backend/model/jwt"
	"testing"
)

func TestToken(t *testing.T) {
	token := jwt.Token{
		Finger:    "8e4539299ad2f7b4c352018a3fac58c1",
		UserAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36",
		Uuid:      "test",
	}
	tokenStr, err := SetToken(token)
	if err != nil {
		t.Error(err)
	}
	claims, err := GetToken(tokenStr)
	if err != nil {
		t.Error(err)
	}
	t.Log(claims)
}
