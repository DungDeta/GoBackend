package manage

import (
	"github.com/gin-gonic/gin"
)

type UserRouter struct{}

func (u *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	// public router
	// private router
	userRouterPrivate := Router.Group("/admin/user")
	{
		userRouterPrivate.GET("/list")
	}
}
