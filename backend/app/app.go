package app

import (
	"github.com/DnsUnlock/Dpanel/backend/api"
	"github.com/DnsUnlock/Dpanel/backend/api/cors"
	"github.com/DnsUnlock/Dpanel/backend/config"
	"github.com/DnsUnlock/Dpanel/backend/utils/log"
	"github.com/gin-gonic/gin"
	"strconv"
)

func Run() {
	r := gin.Default()
	r.Use(cors.Cors())
	api.Router(r.Group("/api"))
	err := r.Run(":" + strconv.Itoa(config.Config.Port))
	if err != nil {
		log.Println(log.ERROR, "服务启动失败", err)
	}
}
