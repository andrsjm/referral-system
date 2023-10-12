package util

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"strings"
)

func GenerateCode(username string) (string, error) {
	username = strings.ToLower(username)

	bytes := make([]byte, 6)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	token := hex.EncodeToString(bytes)

	referralCode := fmt.Sprintf("%s-%s", username, token)

	return referralCode, nil
}
