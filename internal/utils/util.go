package utils

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

func GetUserKey(UserKey string) string {
	return fmt.Sprintf("user:%s", UserKey)
}
func GenerateCliTokenUUID(UserID int) string {
	uuidString := strings.ReplaceAll(uuid.New().String(), "", "")
	return strconv.Itoa(UserID) + "clitoken" + uuidString
}
