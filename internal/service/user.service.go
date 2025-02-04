package service

import (
	"fmt"
	"time"

	"myproject/internal/repo"
	"myproject/internal/utils/crypto"
	"myproject/internal/utils/random"
	"myproject/internal/utils/sendto"
	"myproject/pkg/response"
)

type IUserService interface {
	Register(email string, purpose string) int
}

type userService struct {
	userRepo repo.IUserRepository
	userAuth repo.IUserAuthRepository
}

func (u userService) Register(email string, purpose string) int {
	// TODO implement me
	// 0.Hash Email
	hashEmail := crypto.GetHash(email)
	fmt.Println("Hash Email: ", hashEmail)
	// 1. Check email exist
	if u.userRepo.GetUserByEmail(email) {
		return response.ErrCodeUserExist
	}

	// 2. Create new OTP
	otp := random.GenerateSixDigitOtp()
	if purpose == "test" {
		otp = 123456
	}
	fmt.Printf("Otp: %d\n", otp)
	// 3. Save to Redis with expire time
	err := u.userAuth.AddOTP(hashEmail, otp, int64(10*time.Minute))
	if err != nil {
		return response.ErrInvalidOTP
	}
	// 4. Send email
	err = sendto.SendTextEmail([]string{email}, "test@test.com", otp)
	if err != nil {
		return response.ErrSendEmail
	}
	// 5. Check OTP

	// 6. Save to DB
	return response.ErrCodeSuccess
}

func NewUserService(userRepo repo.IUserRepository, userAuth repo.IUserAuthRepository) IUserService {
	return &userService{
		userRepo: userRepo,
		userAuth: userAuth,
	}
}
