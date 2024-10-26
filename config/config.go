// redis config
package config

import (
	"os"

	"github.com/go-redis/redis"
)

var (
	Rdb *redis.Client
)

func InitRedis() {
	redisHost := os.Getenv("REDIS_HOST")
	if redisHost == "" {
		redisHost = "localhost:" + os.Getenv("REDIS_PORT")
	}
	Rdb = redis.NewClient(&redis.Options{
		Addr:     redisHost + ":" + os.Getenv("REDIS_PORT"),
		Password: "",
		DB:       0,
	})
}
