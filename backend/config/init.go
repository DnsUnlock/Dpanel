package config

import "github.com/DnsUnlock/Dpanel/backend/utils/log"

func init() {
	err := ReadeConfig()
	if err != nil {
		log.Println(log.ERROR, "配置文件读取失败", err)
	}
	log.Println(log.INFO, "配置文件读取成功")
}
