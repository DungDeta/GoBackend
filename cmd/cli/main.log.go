package main

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	encoder := getEncoderLog()
	sync := getWriterSync()
	core := zapcore.NewCore(encoder, sync, zapcore.InfoLevel)
	logger := zap.New(core, zap.AddCaller())

	logger.Info("Info logs ", zap.Int("line", 1))
	logger.Info("Info logs ", zap.Int("line", 2))
}

func getEncoderLog() zapcore.Encoder {
	encoder := zap.NewProductionEncoderConfig()
	encoder.EncodeTime = zapcore.ISO8601TimeEncoder // Định dạng thời gian
	encoder.TimeKey = "Time"
	encoder.EncodeLevel = zapcore.CapitalLevelEncoder // Định dạng level logs thành chữ hoa
	encoder.EncodeCaller = zapcore.ShortCallerEncoder // Hiển thị thông tin file ghi logs
	return zapcore.NewJSONEncoder(encoder)
}

// Ghi logs ra file và console một cách đồng thời
func getWriterSync() zapcore.WriteSyncer {
	file, _ := os.OpenFile("cmd/logs/logs.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	syncFile := zapcore.AddSync(file)
	syncConsole := zapcore.AddSync(os.Stderr)
	return zapcore.NewMultiWriteSyncer(syncFile, syncConsole)
}
