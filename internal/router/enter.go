package router

import (
	"myproject/internal/router/manage"
	"myproject/internal/router/user"
)

type RouterGroup struct {
	User   user.UserRouterGroup
	Manage manage.AdminRouterGroup
}

var RouterGroupApp RouterGroup
