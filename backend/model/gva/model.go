package gva

import "time"

type Model struct {
	ID        uint      `gorm:"primarykey"`         // 主键ID
	CreatedAt time.Time `gorm:"column:created_at;"` // 创建时间
	UpdatedAt time.Time `gorm:"column:updated_at;"` // 更新时间
	DeletedAt *string   `json:"-"`                  // 删除时间
}
