package initialize

import (
	"github.com/gin-gonic/gin"
	"myproject/global"
	"myproject/internal/router"
)

func InitRouter() *gin.Engine {
	var r *gin.Engine
	if global.Config.Server.Mode == "dev" {
		gin.SetMode(gin.DebugMode)
		gin.ForceConsoleColor()
		r = gin.Default()
	} else {
		gin.SetMode(gin.ReleaseMode)
		r = gin.New() // gin.New() không sử dụng Logger và Recovery Middleware
	}
	// Middleware
	r.Use() // Logger
	r.Use() //
	// Router
	manageRouter := router.RouterGroupApp.Manage
	userRouter := router.RouterGroupApp.User
	MainGroup := r.Group("/v1/2024")
	{
		MainGroup.GET("/check_status", func(c *gin.Context) {
			c.JSON(200, "Oke")
		})
	}
	{
		userRouter.InitUserRouter(MainGroup)
		userRouter.InitProductRouter(MainGroup)
	}
	{
		manageRouter.InitAdminRouter(MainGroup)
		manageRouter.InitUserRouter(MainGroup)
	}
	return r
}
