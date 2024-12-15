package db

import (
	"github.com/DnsUnlock/Dpanel/backend/db/redis"
	"github.com/DnsUnlock/Dpanel/backend/db/sql"
	"github.com/DnsUnlock/Dpanel/backend/utils/log"
)

func init() {
	var err error
	err = sql.Init()
	if err != nil {
		log.Println(log.ERROR, "数据库连接失败", err)
	}
	log.Println(log.INFO, "数据库连接成功")
	err = redis.Init()
	if err != nil {
		log.Println(log.ERROR, "Redis连接失败", err)
	}
	log.Println(log.INFO, "Redis连接成功")
}
