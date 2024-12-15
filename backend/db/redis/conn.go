package redis

import (
	"github.com/DnsUnlock/Dpanel/backend/config"
	"github.com/go-redis/redis"
)

func Conn() (db *redis.Client, err error) {
	addr := config.Config.Redis.Host + ":" + config.Config.Redis.Port
	redisDb := redis.NewClient(&redis.Options{
		Addr:         addr,
		Password:     config.Config.Redis.Password,
		DB:           config.Config.Redis.DB,
		PoolSize:     200,
		MinIdleConns: 100,
	})
	_, err = redisDb.Ping().Result()
	if err != nil {
		return nil, err
	}
	return redisDb, nil
}
