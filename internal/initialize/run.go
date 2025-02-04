package initialize

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"myproject/global"
)

func Run() *gin.Engine {
	LoadConfig() // Load Config đầu tiên
	InitLogger() // Khởi tạo Logger để ghi logs sớm nhất
	global.Logger.Info("Logger init success", zap.String("Ok", "Success"))
	InitRedis()
	InitMysqlc()
	InitServiceInterface()
	InitKafka()
	r := InitRouter()
	return r
	// err := r.Run(":8080")
	// if err != nil {
	// 	return nil
	// }
}
