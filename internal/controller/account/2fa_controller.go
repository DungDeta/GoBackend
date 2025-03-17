package account

import (
	"github.com/gin-gonic/gin"
	"myproject/internal/model"
	"myproject/internal/service"
	"myproject/internal/utils/context"
	"myproject/pkg/response"
)

var TwoFa = new(sUser2Fa)

type sUser2Fa struct {
}

// SetupTwoFactorAuth godoc
// @Summary      SetupTwoFactorAuth by user
// @Description	 SetupTwoFactorAuth login by user
// @Tags         accounts 2fa
// @Accept       json
// @Produce      json
// @Param Authorization header string true "Authorization token"
// @Param		 payload body model.SetupTwoFactorAuthInput true "payload"
// @Success      200  {object}  response.ResponseData
// @Failure      500  {object}  response.ErrorResponseData
// @Router       /two-factor/setup [post]
func (f sUser2Fa) SetupTwoFactorAuth(ctx *gin.Context) {
	var params model.SetupTwoFactorAuthInput
	if err := ctx.ShouldBindJSON(&params); err != nil {
		response.ErrorResponse(ctx, response.ErrCodeParamInvalid, err.Error())
		return
	}
	userID, err := context.GetUserIdFromUUID(ctx.Request.Context())
	if err != nil {
		response.ErrorResponse(ctx, response.ErrCodeParamInvalid, err.Error())
		return
	}
	params.UserId = uint32(userID)
	codeStatus, err := service.UserLogin().SetupTwoFactorAuth(ctx.Request.Context(), &params)
	if err != nil {
		response.ErrorResponse(ctx, codeStatus, err.Error())
		return
	}
	response.SuccessResponse(ctx, codeStatus, nil)
}

// VerifyTwoFactorAuth godoc
// @Summary      VerifyTwoFactorAuth by user
// @Description	 VerifyTwoFactorAuth login by user
// @Tags         accounts 2fa
// @Accept       json
// @Produce      json
// @Param Authorization header string true "Authorization token"
// @Param		 payload body model.TwoFactorVerificationInput true "payload"
// @Success      200  {object}  response.ResponseData
// @Failure      500  {object}  response.ErrorResponseData
// @Router       /two-factor/verify [post]
func (f sUser2Fa) VerifyTwoFactorAuth(ctx *gin.Context) {
	var params model.TwoFactorVerificationInput
	if err := ctx.ShouldBindJSON(&params); err != nil {
		response.ErrorResponse(ctx, response.ErrCodeParamInvalid, err.Error())
		return
	}
	userID, err := context.GetUserIdFromUUID(ctx.Request.Context())
	if err != nil {
		response.ErrorResponse(ctx, response.ErrCodeParamInvalid, err.Error())
		return
	}
	params.UserId = uint32(userID)
	codeStatus, err := service.UserLogin().VerifyTwoFactorAuth(ctx.Request.Context(), &params)
	if err != nil {
		response.ErrorResponse(ctx, codeStatus, err.Error())
		return
	}
	response.SuccessResponse(ctx, codeStatus, nil)
}
