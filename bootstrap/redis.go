package bootstrap

import (
	"fmt"

	"github.com/childelins/go-skeleton/pkg/config"
	"github.com/childelins/go-skeleton/pkg/redis"
)

func SetupRedis() {
	// 建立 Redis 连接
	redis.ConnectRedis(
		fmt.Sprintf("%v:%v", config.GetString("redis.host"), config.GetString("redis.port")),
		config.GetString("redis.username"),
		config.GetString("redis.password"),
		config.GetInt("redis.database"),
	)
}
