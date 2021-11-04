// +build dev

package config

import (
	"go.uber.org/zap"
)

var sugarLogger *zap.SugaredLogger

func InitLogger() {
	logger, _ := zap.NewDevelopment()
	defer logger.Sync()
	sugarLogger = logger.Sugar()
}

func Logger() *zap.SugaredLogger {
	return sugarLogger
}
