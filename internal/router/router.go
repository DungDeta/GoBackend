package router

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"myproject/internal/controller"
	"myproject/internal/middlewares"
)

func AA() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Before AA")
		c.Next()
		fmt.Println("Alter AA")
	}
}
func BB() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Before BB")
		c.Next()
		fmt.Println("Alter BB")
	}
}
func CC(c *gin.Context) {
	fmt.Println("Before CC")
	c.Next()
	fmt.Println("Alter CC")
}

func NewRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middlewares.AuthenMiddleware(), BB(), CC)
	v1 := r.Group("/v1")
	{
		v1.GET("/ping/1", controller.NewUserController().UserGetByID)
		v1.GET("/ping", controller.NewPongController().Pong)
		// v1/ping
		// v1.PUT("/ping", pong)
		// v1.POST("/ping", pong)

	}
	err := r.Run()
	if err != nil {
		return nil
	}
	return r
}
