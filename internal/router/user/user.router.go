package user

import (
	"github.com/gin-gonic/gin"
	"myproject/internal/controller/account"
	"myproject/internal/wire"
)

type UserRouter struct {
}

func (r *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	userController, _ := wire.InitUserRouterHandler()

	userRouterPublic := Router.Group("/user")
	{
		userRouterPublic.POST("/register", account.Login.Register)
		userRouterPublic.POST("/login", account.Login.Login)
	}

	userRouterPrivate := Router.Group("/user")
	{
		userRouterPrivate.GET("/get_info", userController.Register)
	}
}
