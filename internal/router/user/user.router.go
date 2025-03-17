package user

import (
	"github.com/gin-gonic/gin"
	"myproject/internal/controller/account"
	"myproject/internal/middlewares"
)

type UserRouter struct {
}

func (r *UserRouter) InitUserRouter(Router *gin.RouterGroup) {

	userRouterPublic := Router.Group("/user")
	{
		userRouterPublic.POST("/register", account.Login.Register)
		userRouterPublic.POST("/login", account.Login.Login)
		userRouterPublic.POST("/verify", account.Login.VerifyOTP)
		userRouterPublic.POST("/updatepassword", account.Login.UpdatePasswordRegister)
	}

	userRouterPrivate := Router.Group("/user")
	userRouterPrivate.Use(middlewares.AuthenMiddleware())
	{
		userRouterPrivate.POST("/two-factor/setup", account.TwoFa.SetupTwoFactorAuth)
		userRouterPrivate.POST("/two-factor/verify", account.TwoFa.VerifyTwoFactorAuth)
	}
}
