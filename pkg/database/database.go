package database

import (
	"database/sql"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	DB    *gorm.DB
	SQLDB *sql.DB
)

// Connect 连接数据库
func Connect(dbConfig gorm.Dialector, logger logger.Interface) {
	var err error

	//使用 gorm.Open 连接数据库
	DB, err = gorm.Open(dbConfig, &gorm.Config{
		// Logger: logger,
	})
	if err != nil {
		panic(err)
	}

	// 获取底层的 sqlDB
	SQLDB, err = DB.DB()
	if err != nil {
		panic(err)
	}

	defer SQLDB.Close()
}
