package redis

import (
	"fmt"
	setting "web_app/settings"

	"github.com/go-redis/redis"
)

var (
	Nil    = redis.Nil
	client *redis.Client
)

func Init(cfg *setting.RedisConfig) (err error) {
	client = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d",
			cfg.Host,
			cfg.Port),
		Password: cfg.Password,
		DB:       cfg.DB,
		PoolSize: cfg.PoolSize,
	})
	_, err = client.Ping().Result()
	return
}

// 封装一个对外的close方法
func Close() {
	_ = client.Close()
}
