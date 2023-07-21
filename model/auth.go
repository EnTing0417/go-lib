package model

import (
	"encoding/pem"
	"fmt"
	"time"

	"crypto/rsa"
	"crypto/x509"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type Token struct {
	ID       string    `json:"id" bson:"id"`
	ExpireAt time.Time `json:"expire_at" bson:"expire_at"`
}

func GenerateToken(claimsMap map[string]interface{}, privateKey *rsa.PrivateKey) (tok string, err error) {
	// Create a new token object
	token := jwt.New(jwt.SigningMethodRS256)

	// Set the claims for the token
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = claimsMap["username"]
	claims["email"] = claimsMap["email"]
	claims["exp"] = claimsMap["exp"]
	claims["user_id"] = claimsMap["user_id"]

	// Generate the token string
	tokenString, err := token.SignedString(privateKey)

	if err != nil {
		fmt.Printf("Error %v", err)
		return "", nil
	}
	return tokenString, err
}

func IsTokenValid(tokenString string, publicKey *rsa.PublicKey) (claims map[string]interface{}, valid bool) {

	// Parse and validate the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Check the signing method
		if token.Method != jwt.SigningMethodRS256 {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		// Return the secret key used for signing
		return publicKey, nil
	})

	if err != nil {
		fmt.Errorf("Token validation failed:", err)
		return nil, false
	}

	// Verify if the token is valid
	if !token.Valid {
		fmt.Errorf("Token is invalid")
		return nil, false
	}

	// Access the token claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		fmt.Errorf("Invalid token claims")
		return nil, false
	}

	expire_at, ok := claims["exp"].(float64)

	if !ok {
		fmt.Errorf("Invalid token expiry")
		return nil, false
	}

	// Extract specific claims
	_claims := map[string]interface{}{
		"username":  claims["username"].(string),
		"email":     claims["email"].(string),
		"user_id":   claims["user_id"].(string),
		"expire_at": expire_at,
	}

	expiredAt := time.Unix(int64(expire_at), 0)

	// Check if the token has expired
	if expiredAt.Before(time.Now()) {
		fmt.Println("Token has expired!")
		return _claims, false
	}
	return _claims, true
}

func ParseRSAPrivateKeyFromConfig(privateKeyPEM string) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode([]byte(privateKeyPEM))

	if block == nil {
		return nil, fmt.Errorf("failed to parse PEM block containing the private key")
	}

	privKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		fmt.Println("Failed to parse RSA private key")
		return nil, err
	}

	rsaPrivateKey, ok := privKey.(*rsa.PrivateKey)
	if !ok {
		fmt.Println("Failed to convert to RSA private key")
		return nil, fmt.Errorf("failed to convert the private key")
	}

	return rsaPrivateKey, nil
}

func ParseRSAPublicKeyFromConfig(publicKeyPEM string) (*rsa.PublicKey, error) {
	block, _ := pem.Decode([]byte(publicKeyPEM))

	if block == nil {
		return nil, fmt.Errorf("failed to parse PEM block containing the public key")
	}

	pubKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	rsaPublicKey, ok := pubKey.(*rsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("key type is not RSA")
	}

	return rsaPublicKey, nil
}

func Hash(password string) (encrypted string, err error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

func VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
