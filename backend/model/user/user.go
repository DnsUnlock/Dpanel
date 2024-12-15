package user

import (
	"github.com/DnsUnlock/Dpanel/backend/model/gva"
	"github.com/DnsUnlock/Dpanel/backend/model/prefix"
)

type User struct {
	gva.Model
	UserName string `json:"user_name" gorm:"type:varchar(255);not null;unique"` // 用户名
	Password string `json:"password" gorm:"type:varchar(255);not null"`         // 密码
	Banned   int    `json:"banned" gorm:"type:int;default:0"`                   // 封禁状态
}

func (User) TableName() string {
	return prefix.Prefix() + "user"
}
