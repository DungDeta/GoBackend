package model

type RegisterInput struct {
	VerifyKey     string `json:"verify_key"` // Mail or Phone
	VerifyType    int    `json:"verify_type"`
	VerifyPurpose string `json:"verify_purpose"`
}
