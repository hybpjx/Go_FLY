package global

import (
	"go.uber.org/zap"
	"gofly/conf"
	"gorm.io/gorm"
)

var (
	// Logger 日志写入 日志句柄
	Logger *zap.SugaredLogger
	// DB    全局调用DB
	DB *gorm.DB
	// RedisClient 全局调用Redis
	RedisClient *conf.RedisClient
)
