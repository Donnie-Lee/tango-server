package redis

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"imserver/config"
	"imserver/middleware/log"
	"time"
)

var RedisPool *redis.Pool

func init() {
	redisConfig := config.Config.RedisConfig
	pool := &redis.Pool{
		MaxIdle:     5,                 // 最大空闲连接数
		IdleTimeout: 240 * time.Second, // 空闲超时时间
		Dial: func() (redis.Conn, error) {
			con, err := redis.Dial("tcp", fmt.Sprintf("%s:%d", redisConfig.Host, redisConfig.Port))
			if err != nil {
				log.Logger.Error("redis connect failed", err)
				panic("redis connect failed" + err.Error())
			}

			if _, authErr := con.Do("AUTH", redisConfig.Password); authErr != nil {
				log.Logger.Error("Redis auth error", authErr)
				panic("Redis auth error" + authErr.Error())
			}
			return con, err
		},
	}

	RedisPool = pool
}
