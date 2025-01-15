package controller

import (
	"github.com/gin-gonic/gin"
	"myproject/internal/service"
	"myproject/pkg/response"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController() *UserController {
	return &UserController{
		userService: service.NewUserService(),
	}
}

// UserGetByID controller -> service -> repo -> models -> db
func (u *UserController) UserGetByID(c *gin.Context) {
	response.SuccessResponse(c, 20001, []string{"user1", "user2"})
}
