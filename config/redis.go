package config

import (
	"flag"
	"log"

	"github.com/garyburd/redigo/redis"
)

var (
	redisAddr = flag.String("addr", "redis://user:@localhost:6379/0", "redis addr")
)

func RedisConnect() *redis.Conn {

	flag.Parse()

	var err error
	redisConn, err := redis.DialURL(*redisAddr)
	if err != nil {
		log.Fatalf("cant connect to redis")
	}

	return &redisConn
}
