package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"myproject/internal/service"
	"myproject/internal/vo"
	"myproject/pkg/response"
)

type UserController struct {
	userService service.IUserService
}

func NewUserController(userService service.IUserService) *UserController {
	return &UserController{
		userService: userService,
	}
}
func (uc *UserController) Register(c *gin.Context) {
	var params vo.UserRegistratorRequest
	if err := c.ShouldBindJSON(&params); err != nil {
		response.ErrorResponse(c, response.ErrCodeParamInvalid, err.Error())
		return
	}
	fmt.Println(params)

	res := uc.userService.Register(params.Email, params.Purpose)
	response.SuccessResponse(c, response.ErrCodeSuccess, res)
}
