package logger

import (
	"os"

	"go.uber.org/zap"
)

var Logger *zap.Logger

func New() *zap.Logger {
	var err error
	
	// 根据环境配置日志级别
	config := zap.NewProductionConfig()
	
	// 开发环境使用更友好的格式
	if os.Getenv("GIN_MODE") != "release" {
		config = zap.NewDevelopmentConfig()
	}
	
	// 自定义输出格式
	config.OutputPaths = []string{"stdout"}
	config.ErrorOutputPaths = []string{"stderr"}
	
	Logger, err = config.Build()
	if err != nil {
		panic(err)
	}
	
	return Logger
}

// Info 记录Info级别日志
func Info(msg string, fields ...zap.Field) {
	Logger.Info(msg, fields...)
}

// Error 记录Error级别日志
func Error(msg string, fields ...zap.Field) {
	Logger.Error(msg, fields...)
}

// Debug 记录Debug级别日志
func Debug(msg string, fields ...zap.Field) {
	Logger.Debug(msg, fields...)
}

// Warn 记录Warn级别日志
func Warn(msg string, fields ...zap.Field) {
	Logger.Warn(msg, fields...)
}

// Fatal 记录Fatal级别日志
func Fatal(msg string, fields ...zap.Field) {
	Logger.Fatal(msg, fields...)
} 