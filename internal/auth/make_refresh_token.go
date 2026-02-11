package auth

import (
	"crypto/rand"
	"encoding/hex"
)

func MakeRefreshToken() (string, error) {
	refreshToken := make([]byte, 32)
	_, err := rand.Read(refreshToken)
	if err != nil {
		return "", err
	}
	encodedRefreshToken := hex.EncodeToString(refreshToken)
	return encodedRefreshToken, nil
}
