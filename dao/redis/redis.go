package redis

import (
	"fmt"
	"go_frame/settings"

	"github.com/go-redis/redis"
)

// 声明一个全局的rdb变量
var rdb *redis.Client

func Init(cfg *settings.RedisConfig) (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d",
			cfg.Host,
			cfg.Port,
		), // redis地址
		Password: cfg.Password, // redis密码，没有则留空
		DB:       cfg.DB,       // 默认数据库，默认是0
		PoolSize: cfg.PoolSize, // 连接池大小
	})

	_, err = rdb.Ping().Result()
	return err
}

func Close() {
	_ = rdb.Close()
}
