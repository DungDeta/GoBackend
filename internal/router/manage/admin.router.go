package manage

import (
	"github.com/gin-gonic/gin"
)

type AdminRouter struct {
}

func (a *AdminRouter) InitAdminRouter(Router *gin.RouterGroup) {
	// public router
	adminPublicRouter := Router.Group("/admin")
	{
		adminPublicRouter.POST("/login")
	}
	// private router
	adminRouterPrivate := Router.Group("/admin")
	{
		adminRouterPrivate.GET("/list")

	}
}
