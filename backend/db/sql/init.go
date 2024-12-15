package sql

import (
	"github.com/DnsUnlock/Dpanel/backend/config"
	"github.com/DnsUnlock/Dpanel/backend/db/sql/mysql"
	"github.com/DnsUnlock/Dpanel/backend/db/sql/sqlite"
	"github.com/DnsUnlock/Dpanel/backend/utils/log"
	"gorm.io/gorm"
)

var Db *gorm.DB

func init() {
	var err error
	switch config.Config.Sql.Driver {
	case "mysql":
		Db, err = mysql.Conn()
	case "sqlite":
		Db, err = sqlite.Conn()
	default:
		log.Println(log.ERROR, "数据库驱动不支持")
	}
	if err != nil {
		log.Println(log.ERROR, "数据库连接失败", err)
	}
	return
}

func Get() *gorm.DB {
	return Db
}
