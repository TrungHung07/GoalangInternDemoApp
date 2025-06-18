package data

import (
	"DemoApp/internal/conf"

	"github.com/go-redis/redis/v8"
)

func NewRedisClient(c *conf.Data_Redis) *redis.Client {
	return redis.NewClient(&redis.Options{
		Network:      c.Network,
		Addr:         c.Addr,
		ReadTimeout:  c.ReadTimeout.AsDuration(),
		WriteTimeout: c.WriteTimeout.AsDuration(),
	})
}

func NewRedisConfig(c *conf.Data) *conf.Data_Redis {
	return c.GetRedis()
}
