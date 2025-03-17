package impl

import (
	"context"
	"database/sql"
	"encoding/json"
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
	"myproject/internal/utils/auth"
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

func (s *sUserLogin) Login(ctx context.Context, in *model.LoginInput) (codeResult int, out model.LoginOutput, err error) {
	userBase, err := s.r.GetOneUserInfo(ctx, in.UserAccount)
	if err != nil {
		return response.ErrCodeLoginFail, out, err
	}
	if !crypto.MatchingPassword(userBase.UserPassword, in.UserPassword, userBase.UserSalt) {
		return response.ErrCodeLoginFail, out, err
	}
	// Check bảo mật 2FA
	isTwoFactorAuth, err := s.r.IsTwoFactorEnabled(ctx, uint32(userBase.UserID))
	if err != nil {
		return response.ErrCodeLoginFail, out, err
	}
	if isTwoFactorAuth > 0 {
		keyUserLoginTwoFactor := crypto.GetHash("2fa:otp" + strconv.Itoa(int(userBase.UserID)))
		err = global.Rdb.Set(ctx, keyUserLoginTwoFactor, "123456", time.Duration(_const.TIME_2FA_OTP_REGISTER)*time.Minute).Err()
		if err != nil {
			return response.ErrCodeTwoFactorAuthFail, out, err
		}
		// Send OTP
		// Get email 2FA
		infoUserTwoFactor, err := s.r.GetTwoFactorMethodByIDAndType(ctx, database.GetTwoFactorMethodByIDAndTypeParams{
			UserID:            uint32(userBase.UserID),
			TwoFactorAuthType: database.PreGoAccUserTwoFactor9999TwoFactorAuthTypeEMAIL,
		})
		if err != nil {
			return response.ErrCodeTwoFactorAuthFail, out, err
		}
		log.Println("send OTP to email: ", infoUserTwoFactor.TwoFactorEmail)
		// Send OTP to email
		// go sendto.SendTextEmail([]string{infoUserTwoFactor.TwoFactorEmail}, _const.HOST_MAIL, "123456")
		out.Message = "Send OTP to email" + infoUserTwoFactor.TwoFactorEmail.String
		return response.ErrCodeSuccess, out, nil
	}
	// Update last login
	go func() {
		err := s.r.LoginUserBase(ctx, database.LoginUserBaseParams{
			UserLoginIp:  sql.NullString{String: "127.0.0.1", Valid: true},
			UserAccount:  in.UserAccount,
			UserPassword: in.UserPassword,
		})
		if err != nil {
			global.Logger.Error("LoginUserBase error: ", zap.Error(err))
		}
	}()
	// Tạo UUID Token
	subToken := utils.GenerateCliTokenUUID(int(userBase.UserID))
	log.Println("subtoken:", subToken)
	// Lấy thông tin user
	infoUser, err := s.r.GetUser(ctx, uint64(userBase.UserID))
	if err != nil {
		return response.ErrCodeLoginFail, out, err
	}
	infoUserJson, err := json.Marshal(infoUser)
	if err != nil {
		return response.ErrCodeLoginFail, out, fmt.Errorf("fail to convert to Json: %w", err)
	}
	// Save infoUserJson vào Redis với key là subToken
	err = global.Rdb.Set(ctx, subToken, infoUserJson, time.Duration(_const.TIME_2FA_OTP_REGISTER)*time.Minute).Err()
	if err != nil {
		return response.ErrCodeLoginFail, out, err
	}
	out.Token, err = auth.CreateToken(subToken)
	if err != nil {
		return
	}
	return response.ErrCodeSuccess, out, nil
}

func (s sUserLogin) Register(ctx context.Context, in *model.RegisterInput) (codeResult int, err error) {
	fmt.Println("VerifyKey: ", in.VerifyKey)
	fmt.Println("VerifyType: ", in.VerifyType)
	hashMail := crypto.GetHash(strings.ToLower(in.VerifyKey))
	fmt.Println("HashMail: ", hashMail)
	userFound, err := s.r.CheckUserBaseExists(ctx, in.VerifyKey)
	if err != nil {
		global.Logger.Error("CheckUserBaseExists error: ", zap.Error(err))
		return response.ErrCodeUserExist, nil
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

func (s *sUserLogin) VerifyOTP(ctx context.Context, in *model.VerifyInput) (out model.VerifyOutput, err error) {
	hashKey := crypto.GetHash(strings.ToLower(in.VerifyKey))
	otpFound, err := global.Rdb.Get(ctx, utils.GetUserKey(hashKey)).Result()
	if err != nil {
		return out, err
	}
	if otpFound != in.VerifyCode {
		return out, fmt.Errorf("OTP khong trung nhau")
	}
	infoOTP, err := s.r.GetInfoOTP(ctx, hashKey)
	if err != nil {
		return out, err
	}
	err = s.r.UpdateUserVerificationStatus(ctx, hashKey)
	if err != nil {
		return out, err
	}
	out.Token = infoOTP.VerifyKeyHash
	out.Message = "Success"
	return out, err
}
func (s *sUserLogin) UpdatePasswordRegister(ctx context.Context, token string, password string) (userId int64, err error) {
	infoOTP, err := s.r.GetInfoOTP(ctx, token)
	if err != nil {
		return response.ErrNotVerify, err
	}
	if infoOTP.IsVerified.Int32 == 0 {
		return response.ErrCodeOtpNotExisted, err
	}
	userBase := database.AddUserBaseParams{}
	userBase.UserAccount = infoOTP.VerifyKey
	userBase.UserSalt, err = crypto.GenSalt(16)
	if err != nil {
		return response.ErrCodeOtpNotExisted, err
	}
	userBase.UserPassword = crypto.HashPassword(password, userBase.UserSalt)
	newUserBase, err := s.r.AddUserBase(ctx, userBase)
	if err != nil {
		return response.ErrCodeOtpNotExisted, err
	}
	userId, err = newUserBase.LastInsertId()
	if err != nil {
		return response.ErrCodeOtpNotExisted, err
	}
	newUserInfo, err := s.r.AddUserHaveUserId(ctx, database.AddUserHaveUserIdParams{
		UserID:       uint64(userId),
		UserAccount:  infoOTP.VerifyKey,
		UserNickname: sql.NullString{String: infoOTP.VerifyKey, Valid: true},
		UserAvatar:   sql.NullString{String: "", Valid: true},
		UserState:    1,
		UserMobile:   sql.NullString{String: "", Valid: true},
		UserGender: sql.NullInt16{
			Int16: 0,
			Valid: true,
		},
		UserBirthday:         sql.NullTime{Time: time.Time{}, Valid: false},
		UserEmail:            sql.NullString{String: infoOTP.VerifyKey, Valid: true},
		UserIsAuthentication: 1,
	})
	if err != nil {
		return response.ErrCodeOtpNotExisted, err
	}
	userId, err = newUserInfo.LastInsertId()
	if err != nil {
		return response.ErrCodeOtpNotExisted, err
	}
	return int64(userId), err
}
func (s *sUserLogin) SetupTwoFactorAuth(ctx context.Context, in *model.SetupTwoFactorAuthInput) (codeResult int, err error) {
	isTwoFactorAuth, err := s.r.IsTwoFactorEnabled(ctx, in.UserId)
	if err != nil {
		return response.ErrCodeTwoFactorAuthSetupFail, err
	}
	if isTwoFactorAuth > 0 {
		return response.ErrCodeTwoFactorAuthSetupFail, fmt.Errorf("two factor auth is already enabled")
	}
	err = s.r.EnableTwoFactorTypeEmail(ctx, database.EnableTwoFactorTypeEmailParams{
		UserID:            in.UserId,
		TwoFactorAuthType: database.PreGoAccUserTwoFactor9999TwoFactorAuthTypeEMAIL,
		TwoFactorEmail:    sql.NullString{String: in.TwoFactorEmail, Valid: true},
	})
	if err != nil {
		return response.ErrCodeTwoFactorAuthSetupFail, err
	}
	// Send OTP
	keyUserTwoFactor := crypto.GetHash("2fa" + strconv.Itoa(int(in.UserId)))
	err = global.Rdb.Set(ctx, keyUserTwoFactor, 123456, time.Duration(_const.TIME_OTP_REGISTER)*time.Minute).Err()
	// Dùng goroutine cũng được
	if err != nil {
		return response.ErrCodeTwoFactorAuthSetupFail, err
	}
	return 200, nil
}

func (s *sUserLogin) VerifyTwoFactorAuth(ctx context.Context, in *model.TwoFactorVerificationInput) (codeResult int, err error) {
	isTwoFactorAuth, err := s.r.IsTwoFactorEnabled(ctx, in.UserId)
	if err != nil {
		return response.ErrCodeTwoFactorAuthFail, err
	}
	if isTwoFactorAuth > 0 {
		return response.ErrCodeTwoFactorAuthFail, fmt.Errorf("two factor auth is not enabled")
	}
	// Check Redis
	keyUserTwoFactor := crypto.GetHash("2fa" + strconv.Itoa(int(in.UserId)))
	otpFound, err := global.Rdb.Get(ctx, keyUserTwoFactor).Result()
	if errors.Is(err, redis.Nil) {
		return response.ErrCodeTwoFactorAuthFail, fmt.Errorf("OTP not found")
	} else if err != nil {
		return response.ErrCodeTwoFactorAuthFail, err
	}
	// Check OTP
	if otpFound != in.TwoFactorCode {
		return response.ErrCodeTwoFactorAuthFail, fmt.Errorf("OTP not match")
	}
	// Update Two Factor Auth
	err = s.r.UpdateTwoFactorStatus(ctx, database.UpdateTwoFactorStatusParams{
		UserID:            in.UserId,
		TwoFactorAuthType: database.PreGoAccUserTwoFactor9999TwoFactorAuthTypeEMAIL,
	})
	if err != nil {
		return response.ErrCodeTwoFactorAuthFail, err
	}
	// Xóa OTP
	err = global.Rdb.Del(ctx, keyUserTwoFactor).Err()
	if err != nil {
		return response.ErrCodeTwoFactorAuthFail, fmt.Errorf("delete OTP fail")
	}
	return response.ErrCodeSuccess, nil
}
