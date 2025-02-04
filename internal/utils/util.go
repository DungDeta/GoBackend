package utils

import (
	"fmt"
)

func GetUserKey(UserKey string) string {
	return fmt.Sprintf("user:%s", UserKey)
}
