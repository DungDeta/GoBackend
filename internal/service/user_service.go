package service

import (
	"context"

	"myproject/global"
	"myproject/internal/model"
)

type (
	IUserLogin interface {
		Login(ctx context.Context) error
		Register(ctx context.Context, in *model.RegisterInput) (codeResult int, err error)
		VerifyOTP(ctx context.Context) error
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
		global.Logger.Error("mplement localUserAdmin is not found for interface IUserAdmin")
		panic("implement localUserAdmin is not found for interface IUserAdmin")
	}
	return localUserAdmin
}
func InitUserAdmin(i IUserAdmin) {
	localUserAdmin = i
}

func UserInfo() IUserInfo {
	if localUserInfo == nil {
		global.Logger.Error("mplement localUserAdmin is not found for interface IUserAdmin")
		panic("implement localUserAdmin is not found for interface IUserAdmin")
	}
	return localUserInfo
}
func InitUserInfo(i IUserInfo) {
	localUserInfo = i
}
func UserLogin() IUserLogin {
	if localUserLogin == nil {
		global.Logger.Error("mplement localUserAdmin is not found for interface IUserAdmin")
		panic("implement localUserAdmin is not found for interface IUserAdmin")
	}
	return localUserLogin
}
func InitUserLogin(i IUserLogin) {
	localUserLogin = i
}
