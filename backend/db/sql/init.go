package sql

import (
	"github.com/DnsUnlock/Dpanel/backend/config"
	"github.com/DnsUnlock/Dpanel/backend/db/sql/mysql"
	"github.com/DnsUnlock/Dpanel/backend/db/sql/sqlite"
	"gorm.io/gorm"
)

var Db *gorm.DB

func Init() (err error) {
	switch config.Config.Sql.Driver {
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
