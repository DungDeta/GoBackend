//go:build wireinject

package wire

import (
	"github.com/google/wire"
	"myproject/internal/controller"
	"myproject/internal/repo"
	"myproject/internal/service"
)

func InitUserRouterHandler() (*controller.UserController, error) {
	wire.Build(
		repo.NewUserRepository,
		repo.NewUserAuthRepository,
		service.NewUserService,
		controller.NewUserController,
	)
	return new(controller.UserController), nil

}
