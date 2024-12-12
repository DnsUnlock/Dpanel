package mysql

import (
	"github.com/DnsUnlock/Dpanel/backend/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
)

func Conn() (dB *gorm.DB, err error) {
	dB, err = gorm.Open(
		mysql.New(
			mysql.Config{
				DSN:                       config.Config.Database.Connection,
				SkipInitializeWithVersion: false,
			},
		),
		&gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true, // 禁用表名复数
			},
		},
	)
	if err != nil {
		return
	}
	sqlDB, err := dB.DB()
	if err != nil {
		return
	}
	sqlDB.SetMaxIdleConns(config.Config.Database.MaxIdleCons)
	sqlDB.SetMaxOpenConns(config.Config.Database.MaxOpenCons)
	sqlDB.SetConnMaxLifetime(time.Duration(config.Config.Database.MaxLifeTime) * time.Second)
	return dB, err
}
