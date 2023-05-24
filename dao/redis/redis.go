package redis

import (
	"fmt"

	"github.com/go-redis/redis"
	"github.com/spf13/viper"
)

// 声明一个全局的rdb变量
var rdb *redis.Client

func Init() (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d",
			viper.GetString("redis.host"),
			viper.GetInt("redis.host"),
		), // redis地址
		Password: viper.GetString("redis.password"), // redis密码，没有则留空
		DB:       viper.GetInt("redis.db"),          // 默认数据库，默认是0
		PoolSize: viper.GetInt("redis.pool_size"),   // 连接池大小
	})

	_, err = rdb.Ping().Result()
	return err
}

func Close() {
	_ = rdb.Close()
}
