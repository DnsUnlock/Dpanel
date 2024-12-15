package user

import "github.com/DnsUnlock/Dpanel/backend/model/gva"

type Email struct {
	gva.Model
	UserId uint   `json:"user_id" gorm:"not null;index"`        // 用户ID，不能为空，并添加索引
	Mail   string `json:"mail" gorm:"type:varchar(255);unique"` // 邮箱地址，字段类型为varchar(255)，并确保唯一
}
