package initialize

import (
	"myproject/global"
	"myproject/internal/database"
	"myproject/internal/service"
	"myproject/internal/service/impl"
)

func InitServiceInterface() {
	queries := database.New(global.Mdbc)
	service.InitUserLogin(impl.NewUserLoginImpl(queries))
}
