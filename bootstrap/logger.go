package bootstrap

import (
	"github.com/childelins/go-skeleton/pkg/config"
	"github.com/childelins/go-skeleton/pkg/log"
)

// SetupLogger 初始化 Logger
func SetupLogger() {
	log.InitLogger(
		log.WithFilename(config.GetString("log.filename")),
		log.WithMaxSize(config.GetInt("log.max_size")),
		log.WithMaxBackup(config.GetInt("log.max_backup")),
		log.WithMaxAge(config.GetInt("log.max_age")),
		log.WithCompress(config.GetBool("log.compress")),
		log.WithLogType(config.GetString("log.type")),
		log.WithLevel(config.GetString("log.level")),
	)
}
