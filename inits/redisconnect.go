package inits

import (
	"os"

	"github.com/go-redis/redis"
)

var Redis *redis.Client

func ConnectRedis() {
	var err error
	Redis = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_DB_HOST"),
		Password: os.Getenv("REDIS_DB_PASSWORD"),
		DB:       0,
	})
	_, err = Redis.Ping().Result()
	if err != nil {
		panic(err)
	}
}
