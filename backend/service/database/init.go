package database

import (
	"github.com/DnsUnlock/Dpanel/backend/config"
	"github.com/DnsUnlock/Dpanel/backend/service/database/mysql"
	"github.com/DnsUnlock/Dpanel/backend/service/database/sqlite"
	"gorm.io/gorm"
)

var Db *gorm.DB

func Init() (err error) {
	switch config.Config.Database.Driver {
	case "mysql":
		Db, err = mysql.Conn()
	case "sqlite":
		Db, err = sqlite.Conn()
	}
	return
}

func Get() *gorm.DB {
	return Db
}
