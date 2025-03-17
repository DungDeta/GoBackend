package initialize

import (
	"github.com/gin-gonic/gin"
	"myproject/global"
	"myproject/internal/middlewares"
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
	// r.Use() // Logger
	// r.Use() //
	// Router
	r.Use(middlewares.NewRateLimiter().GlobalRateLimiter())
	r.GET("/ping/100", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Use(middlewares.NewRateLimiter().PublicAPIRateLimiter())
	r.GET("/ping/80", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Use(middlewares.NewRateLimiter().UserAndPrivateRateLimiter())
	r.GET("/ping/50", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
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
		userRouter.InitTicketRouter(MainGroup)
	}
	{
		manageRouter.InitAdminRouter(MainGroup)
		manageRouter.InitUserRouter(MainGroup)
	}
	return r
}
