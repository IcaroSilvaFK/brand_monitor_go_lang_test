package configs

import "github.com/go-redis/redis/v8"

type RedisConfigs struct {
	Redis *redis.Client
}

func NewRedisConfigs() *RedisConfigs {

	config := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	return &RedisConfigs{
		Redis: config,
	}
}
