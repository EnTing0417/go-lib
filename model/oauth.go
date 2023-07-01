package model

import (
	"crypto/rand"
	"encoding/base64"
)

func GenerateAntiForgeryStateToken() (string, error) {
	tokenBytes := make([]byte, 32) 
	_, err := rand.Read(tokenBytes)
	if err != nil {
		return "", err
	}
	token := base64.RawURLEncoding.EncodeToString(tokenBytes)

	return token, nil
}
