package initRedis

import (
	"fatfox-single-sign-on/global"
	"fatfox-single-sign-on/init/initViper"
	"fmt"
	"github.com/go-redis/redis/v8"
)

func InitRedis(config *initViper.Config) {
	// 初始化数据库
	global.RedisDb = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", config.REDIS.Host, config.REDIS.Port),
		Password: config.REDIS.Password,
		DB:       0,
	})
}
