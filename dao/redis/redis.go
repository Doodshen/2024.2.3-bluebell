package redis

import (
	"fmt"
	setting "web_app/settings"

	"github.com/go-redis/redis"
)

var rdb *redis.Client

func Init(cfg *setting.RedisConfig) (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d",
			cfg.Host,
			cfg.Port),
		Password: cfg.Password,
		DB:       cfg.DB,
		PoolSize: cfg.PoolSize,
	})
	_, err = rdb.Ping().Result()
	return
}

// 封装一个对外的close方法
func Close() {
	_ = rdb.Close()
}
