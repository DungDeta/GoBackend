package main

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "myproject/cmd/swag/docs"
	"myproject/internal/initialize"
)

var pingCounter = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "ping_counter",
		Help: "The total number of ping",
	})

func ping(c *gin.Context) {
	pingCounter.Inc()
	c.JSON(200, gin.H{
		"message": "ok",
	})
}

// @title           API Docs Ecommerce Backend
// @version         1.0.0
// @description     This is a sample server celler server.
// @termsOfService	https://github.com/DungDeta/GoBackend

// @contact.name	Deta
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /v1/2024
func main() {
	r := initialize.Run()
	prometheus.MustRegister(pingCounter)
	r.GET("/ping", ping)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))
	err := r.Run(":8080")
	if err != nil {
		return
	}
}
