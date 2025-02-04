package main

import (
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"myproject/internal/initialize"
)

// @title           API Docs Ecommerce Backend
// @version         1.0.0
// @description     This is a sample server celler server.
// @termsOfService

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1
func main() {
	r := initialize.Run()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	err := r.Run(":8080")
	if err != nil {
		return
	}
}
