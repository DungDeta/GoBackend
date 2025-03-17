package service

import (
	"context"

	"myproject/global"
	"myproject/internal/model"
)

type (
	IUserLogin interface {
		Login(ctx context.Context, in *model.LoginInput) (codeResult int, out model.LoginOutput, err error)
		Register(ctx context.Context, in *model.RegisterInput) (codeResult int, err error)
		VerifyOTP(ctx context.Context, in *model.VerifyInput) (out model.VerifyOutput, err error)
		UpdatePasswordRegister(ctx context.Context, token string, password string) (userId int64, err error)
		SetupTwoFactorAuth(ctx context.Context, in *model.SetupTwoFactorAuthInput) (codeResult int, err error)
		VerifyTwoFactorAuth(ctx context.Context, in *model.TwoFactorVerificationInput) (codeResult int, err error)
	}
	IUserInfo interface {
		GetUserInfo(ctx context.Context) error
		GetAllUserInfo(ctx context.Context) error
	}
	IUserAdmin interface {
		RemoveUser(ctx context.Context) error
		FindOneUser(ctx context.Context) error
	}
)

var (
	localUserLogin IUserLogin
	localUserInfo  IUserInfo
	localUserAdmin IUserAdmin
)

func UserAdmin() IUserAdmin {
	if localUserAdmin == nil {
		global.Logger.Error("implement localUserAdmin is not found for interface IUserAdmin")
		panic("implement localUserAdmin is not found for interface IUserAdmin")
	}
	return localUserAdmin
}
func InitUserAdmin(i IUserAdmin) {
	localUserAdmin = i
}

func UserInfo() IUserInfo {
	if localUserInfo == nil {
		global.Logger.Error("implement localUserAdmin is not found for interface IUserAdmin")
		panic("implement localUserAdmin is not found for interface IUserAdmin")
	}
	return localUserInfo
}
func InitUserInfo(i IUserInfo) {
	localUserInfo = i
}
func UserLogin() IUserLogin {
	if localUserLogin == nil {
		global.Logger.Error("implement localUserAdmin is not found for interface IUserAdmin")
		panic("implement localUserAdmin is not found for interface IUserAdmin")
	}
	return localUserLogin
}
func InitUserLogin(i IUserLogin) {
	localUserLogin = i
}
