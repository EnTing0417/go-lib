package model

import (
    "log"
	"time"
	"fmt"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"io"
)

func IsTokenValid(token *AccessToken) (bool){
	// Check if the token has expired
	if token.Expiry.Before(time.Now()) {
		fmt.Println("Token has expired!")
		return false
	} 
	fmt.Println("Token is still valid.")
	return true
}

func EncryptToken(token string) (encryptedString string, err error){

	key := []byte("0123456789abcdef0123456789abcdef")

	nonce := make([]byte, 12)
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		log.Fatal(err)
	}

	// Create a new AES cipher block using the encryption key
	block, err := aes.NewCipher(key)
	if err != nil {
		log.Fatal(err)
		return "",err
	}

	// Create a GCM (Galois Counter Mode) cipher
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		log.Fatal(err)
		return "",err
	}

	// Encrypt the login token
	_encryptedToken := aesgcm.Seal(nil, nonce, []byte(token), nil)

	// Combine the IV and encrypted token into a single byte slice
	ciphertext := append(nonce, _encryptedToken...)

	// Encode the ciphertext as a base64 string
	encryptedString = base64.StdEncoding.EncodeToString(ciphertext)

	return encryptedString,nil
}