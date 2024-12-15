package redis

import "github.com/go-redis/redis"

var Rdb *redis.Client

func Init() (err error) {
	Rdb, err = Conn()
	return
}

func Get() *redis.Client {
	return Rdb
}
