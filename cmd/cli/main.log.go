package main

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	// suger := zap.NewExample().Sugar()
	// suger.Infof("Hello %s, World!", "Zap")

	// logger := zap.NewExample()
	// logger.Info("Hello, World!")
	//
	// logger, _ = zap.NewDevelopment()
	// logger.Info("Hello, World!") // Có thêm thông tin về file ghi log và thời gian
	//
	// logger, _ = zap.NewProduction()
	// logger.Info("Hello, World!") // có thông tin về timestamp và level log

	encoder := getEncoderLog()
	sync := getWriterSync()
	core := zapcore.NewCore(encoder, sync, zapcore.InfoLevel)
	logger := zap.New(core, zap.AddCaller())

	logger.Info("Info log ", zap.Int("line", 1))
	logger.Info("Info log ", zap.Int("line", 2))
}

func getEncoderLog() zapcore.Encoder {
	encoder := zap.NewProductionEncoderConfig()
	encoder.EncodeTime = zapcore.ISO8601TimeEncoder // Định dạng thời gian
	encoder.TimeKey = "Time"
	encoder.EncodeLevel = zapcore.CapitalLevelEncoder // Định dạng level log thành chữ hoa
	encoder.EncodeCaller = zapcore.ShortCallerEncoder // Hiển thị thông tin file ghi log
	return zapcore.NewJSONEncoder(encoder)
}

// Ghi log ra file và console một cách đồng thời
func getWriterSync() zapcore.WriteSyncer {
	file, _ := os.OpenFile("cmd/log/log.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	syncFile := zapcore.AddSync(file)
	syncConsole := zapcore.AddSync(os.Stderr)
	return zapcore.NewMultiWriteSyncer(syncFile, syncConsole)
}
