package inits

import (
	"os"

	"github.com/go-redis/redis"
)

var RedisDB *redis.Client

func init() {
	var err error
	RedisDB := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: os.Getenv("REDIS_DB"),
		DB:       0,
	})
	_, err = RedisDB.Ping().Result()
	if err != nil {
		panic(err)
	}
}
