package repo

import (
	"fmt"
	"time"

	"myproject/global"
)

type IUserAuthRepository interface {
	AddOTP(email string, otp int, expireTime int64) error
}

type userAuthRepository struct {
}

func (u *userAuthRepository) AddOTP(email string, otp int, expireTime int64) error {
	// TODO implement me
	key := fmt.Sprintf("usr:%s:otp", email)                                 // usr:mail@mail.com:otp
	return global.Rdb.SetEx(ctx, key, otp, time.Duration(expireTime)).Err() // Add OTP to Redis
}

func NewUserAuthRepository() IUserAuthRepository {
	return &userAuthRepository{}

}
