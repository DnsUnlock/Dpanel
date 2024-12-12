package main

import (
	"github.com/DnsUnlock/Dpanel/backend/api"
	"github.com/DnsUnlock/Dpanel/backend/api/cors"
	"github.com/DnsUnlock/Dpanel/backend/config"
	"github.com/DnsUnlock/Dpanel/backend/service/database"
	"github.com/DnsUnlock/Dpanel/backend/utils/log"
	"github.com/gin-gonic/gin"
	"strconv"
)

func init() {
	err := config.Init()
	if err != nil {
		log.Println(log.ERROR, "配置文件读取失败", err)
	}
	err = database.Init()
	if err != nil {
		log.Println(log.ERROR, "数据库连接失败", err)
	}
}

func main() {
	r := gin.Default()
	r.Use(cors.Cors())
	api.Router(r.Group("/api"))
	err := r.Run(":" + strconv.Itoa(config.Config.Port))
	if err != nil {
		log.Println(log.ERROR, "服务启动失败", err)
	}
}
