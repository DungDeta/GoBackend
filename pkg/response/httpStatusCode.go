package response

const (
	ErrCodeSuccess      = 20001 // Thành công
	ErrCodeParamInvalid = 20003 // Lỗi email không hợp lệ

	ErrTokenInvalid = 30001
)

var msg = map[int]string{
	ErrCodeSuccess:      "Success",
	ErrCodeParamInvalid: "Email is invalid",
	ErrTokenInvalid:     "Token is invalid",
}
