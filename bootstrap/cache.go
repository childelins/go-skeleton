package bootstrap

import (
	"fmt"

	"github.com/childelins/go-skeleton/pkg/cache"
	"github.com/childelins/go-skeleton/pkg/config"
)

// SetupCache 缓存
func SetupCache() {
	// 初始化缓存专用的 redis client, 使用专属缓存 DB
	rds := cache.NewRedisStore(
		fmt.Sprintf("%v:%v", config.GetString("redis.host"), config.GetString("redis.port")),
		config.GetString("redis.username"),
		config.GetString("redis.password"),
		config.GetInt("redis.database_cache"),
	)

	cache.InitWithCacheStore(rds)
}
