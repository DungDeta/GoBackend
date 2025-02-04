package account

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"myproject/global"
	"myproject/internal/model"
	"myproject/internal/service"
	"myproject/pkg/response"
)

var Login = new(cUserLogin)

type cUserLogin struct {
}

func (c *cUserLogin) Login(ctx *gin.Context) {
	err := service.UserLogin().Login(ctx)
	if err != nil {
		response.ErrorResponse(ctx, response.ErrCodeParamInvalid, err.Error())
	}
	response.SuccessResponse(ctx, response.ErrCodeSuccess, nil)
}
func (c cUserLogin) Register(ctx *gin.Context) {
	var params model.RegisterInput
	if err := ctx.ShouldBindJSON(&params); err != nil {
		response.ErrorResponse(ctx, response.ErrCodeParamInvalid, err)
		return
	}
	codeStatus, err := service.UserLogin().Register(ctx, &params)
	if err != nil {
		global.Logger.Error("Error Reg User OTP", zap.Error(err))
		response.ErrorResponse(ctx, codeStatus, err)
		return
	}
	response.SuccessResponse(ctx, response.ErrCodeSuccess, nil)
}
