package initialize

import (
	"myproject/global"
	"myproject/pkg/logger"
)

func InitLogger() {
	global.Logger = logger.NewLoggerZap(global.Config.Logger)
}
