package impl

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"myproject/global"
	_const "myproject/internal/const"
	"myproject/internal/database"
	"myproject/internal/model"
	"myproject/internal/utils"
	"myproject/internal/utils/crypto"
	"myproject/internal/utils/random"
	"myproject/pkg/response"
)

type sUserLogin struct {
	r *database.Queries
}

func NewUserLoginImpl(r *database.Queries) *sUserLogin {
	return &sUserLogin{
		r: r,
	}
}

func (s *sUserLogin) Login(ctx context.Context) error {
	return nil
}

func (s sUserLogin) Register(ctx context.Context, in *model.RegisterInput) (codeResult int, err error) {
	fmt.Println("VerifyKey: ", in.VerifyKey)
	fmt.Println("VerifyType: ", in.VerifyType)
	hashMail := crypto.GetHash(strings.ToLower(in.VerifyKey))
	fmt.Println("HashMail: ", hashMail)
	userFound, err := s.r.CheckUserBaseExists(ctx, in.VerifyKey)
	if err != nil {
		global.Logger.Error("CheckUserBaseExists error: ", zap.Error(err))
		return 0, err
	}
	if userFound > 0 {
		return response.ErrCodeUserExist, nil
	}
	// Create OTP and save to Redis
	userKey := utils.GetUserKey(hashMail)
	otpFound, err := global.Rdb.Get(ctx, userKey).Result()
	switch {
	case errors.Is(err, redis.Nil):
		fmt.Println("OTP not found")
	case err != nil:
		global.Logger.Error("Get OTP error: ", zap.Error(err))
	case otpFound != "":
		return response.ErrCodeUserExist, nil
	}
	otp := random.GenerateSixDigitOtp()
	if in.VerifyPurpose == "test" {
		otp = 123456
	}
	fmt.Println("OTP: ", otp)
	err = global.Rdb.Set(ctx, userKey, otp, time.Duration(_const.TIME_OTP_REGISTER)*time.Minute).Err()
	if err != nil {
		global.Logger.Error("Save OTP to Redis error: ", zap.Error(err))
		return response.ErrInvalidOTP, err
	}
	switch in.VerifyType {
	case _const.EMAIL:
		// err = sendto.SendTextEmail([]string{in.VerifyKey}, _const.HOST_MAIL, otp)
		// err = nil
		// if err != nil {
		// 	return response.ErrSendEmail, err
		// }
		// 7. save OTP to MYSQL
		result, err := s.r.InsertOTPVerify(ctx, database.InsertOTPVerifyParams{
			VerifyOtp:     strconv.Itoa(otp),
			VerifyType:    sql.NullInt32{Int32: 1, Valid: true},
			VerifyKey:     in.VerifyKey,
			VerifyKeyHash: hashMail,
		})

		if err != nil {
			return response.ErrSendEmail, err
		}

		// 8. getlasId
		lastIdVerifyUser, err := result.LastInsertId()
		if err != nil {
			return response.ErrSendEmail, err
		}
		log.Println("lastIdVerifyUser", lastIdVerifyUser)
		return response.ErrCodeSuccess, nil
	case _const.PHONE:
		return response.ErrCodeSuccess, nil
	default:
		return response.ErrCodeSuccess, nil
	}
}

func (s *sUserLogin) VerifyOTP(ctx context.Context) error {
	return nil
}
