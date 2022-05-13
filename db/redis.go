package db

import (
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
)

var RedisClient *redis.Client

func init() {
    viper.SetConfigFile("../DEV.env")
    viper.ReadInConfig()
    redisURL := viper.GetString("REDIS_URL")
    RedisClient = redis.NewClient(&redis.Options {
        Addr: redisURL,
    })
}
