package inits

import (
	"os"

	"github.com/go-redis/redis"
)

var Redis *redis.Client

func ConnecRedis() {
	var err error
	Redis = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: os.Getenv("REDIS_DB"),
		DB:       0,
	})
	_, err = Redis.Ping().Result()
	if err != nil {
		panic(err)
	}
}
