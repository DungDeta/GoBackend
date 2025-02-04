package logger

import (
	"os"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"myproject/pkg/setting"
)

type LoggerZap struct {
	*zap.Logger
}

func NewLoggerZap(config setting.LoggerSetting) *LoggerZap {
	loglevel := config.LogLevel
	// debug -> info -> warn -> error -> dpanic -> panic -> fatal
	var level zapcore.Level
	switch loglevel {
	case "debug":
		level = zapcore.DebugLevel
	case "info":
		level = zapcore.InfoLevel
	case "warn":
		level = zapcore.WarnLevel
	case "error":
		level = zapcore.ErrorLevel
	case "dpanic":
		level = zapcore.DPanicLevel
	case "panic":
		level = zapcore.PanicLevel
	}
	encoder := getEncoderLog()
	hook := lumberjack.Logger{
		Filename:   config.FileLogName, // Đường dẫn file logs
		MaxSize:    config.MaxSize,     // megabytes
		MaxBackups: config.MaxBackups,
		MaxAge:     config.MaxAge,   // days
		Compress:   config.Compress, // disabled by default
	}
	core := zapcore.NewCore(
		encoder,
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook)),
		level)
	return &LoggerZap{zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))}
}

func getEncoderLog() zapcore.Encoder {
	encoder := zap.NewProductionEncoderConfig()
	encoder.EncodeTime = zapcore.ISO8601TimeEncoder // Định dạng thời gian
	encoder.TimeKey = "Time"
	encoder.EncodeLevel = zapcore.CapitalLevelEncoder // Định dạng level logs thành chữ hoa
	encoder.EncodeCaller = zapcore.ShortCallerEncoder // Hiển thị thông tin file ghi logs
	return zapcore.NewJSONEncoder(encoder)
}
