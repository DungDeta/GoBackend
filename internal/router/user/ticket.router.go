package user

import (
	"github.com/gin-gonic/gin"
	"myproject/internal/controller/ticket"
)

type TicketRouter struct {
}

func (r *TicketRouter) InitTicketRouter(Router *gin.RouterGroup) {
	// public router
	TicketRouterPublic := Router.Group("/ticket")
	{
		TicketRouterPublic.GET("/item/:id", ticket.TicketItem.GetTicketDetailById)
	}
	// private router
}
