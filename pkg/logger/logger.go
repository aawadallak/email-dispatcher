//go:build !dev

package logger

import (
	"fmt"
	"time"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Zap *zap.SugaredLogger

func InitLogger() {
	writerSyncer := getLogWriter()

	encoder := getEncoder()

	core := zapcore.NewCore(encoder, writerSyncer, zapcore.DebugLevel)

	logger := zap.New(core, zap.AddCaller())

	Zap = logger.Sugar()
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getLogWriter() zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   fmt.Sprintf(`./logs/%s.log`, time.Now().Format("02-01-2006")),
		MaxSize:    10,
		MaxAge:     30,
		MaxBackups: 5,
		Compress:   false,
	}
	return zapcore.AddSync(lumberJackLogger)
}

func Instance() *zap.SugaredLogger {
	return Zap
}
