package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

var (
	Sugar  *zap.SugaredLogger
	Logger *zap.Logger
)

// Create 创建日志
func Create() {
	writer := zapcore.AddSync(os.Stdout)

	// 格式相关的配置
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder   // 修改时间戳的格式
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder // 日志级别使用大写显示
	encoder := zapcore.NewConsoleEncoder(encoderConfig)

	core := zapcore.NewCore(encoder, writer, zapcore.InfoLevel)
	Logger = zap.New(core, zap.AddCaller()) // 增加 caller 信息
	Sugar = Logger.Sugar()
}
