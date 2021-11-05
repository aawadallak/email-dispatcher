//go:build dev

package logger

import (
	"go.uber.org/zap"
)

var Zap *zap.SugaredLogger

func InitLogger() {
	logger, _ := zap.NewDevelopment()
	defer logger.Sync()
	Zap = logger.Sugar()
}

func Instance() *zap.SugaredLogger {
	return Zap
}
