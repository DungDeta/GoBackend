package response

const (
	ErrCodeSuccess      = 20001 // Thành công
	ErrCodeParamInvalid = 20003 // Lỗi email không hợp lệ

	ErrTokenInvalid = 30001
	ErrInvalidOTP   = 30002

	ErrCodeUserExist = 40001
	ErrSendEmail     = 50001

	ErrCodeOtpNotExisted = 60001
)

var msg = map[int]string{
	ErrCodeSuccess:      "Success",
	ErrCodeParamInvalid: "Email is invalid",
	ErrTokenInvalid:     "Token is invalid",
	ErrInvalidOTP:       "OTP is invalid",
	ErrCodeUserExist:    "User is exist",
	ErrSendEmail:        "Send email error",
}
