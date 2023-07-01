package model

import (
    "time"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"github.com/dgrijalva/jwt-go"

)

type Token struct {
	ID string `json:"id" bson:"id"`
    ExpireAt time.Time `json:"expire_at" bson:"expire_at"`
}

func GenerateToken(claimsMap map[string]interface{}, secret string) (tok string,err error){
	// Create a new token object
	token := jwt.New(jwt.SigningMethodHS256)

	// Set the claims for the token
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = claimsMap["username"]
	claims["email"] = claimsMap["email"]
	claims["exp"] = claimsMap["exp"] 

	// Generate the token string
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		fmt.Errorf("Error %v", err)
		return "", nil
	}
	return tokenString, err
}

func IsTokenValid(tokenString string, secretKey string) ( claims map[string]interface{},valid bool) {

	// Parse and validate the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Check the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		// Return the secret key used for signing
		return []byte(secretKey), nil
	})
	if err != nil {
		fmt.Errorf("Token validation failed:", err)
		return nil,false
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
		return nil,false
	}

	expire_at, ok := claims["exp"].(float64)
	
	if !ok {
		fmt.Errorf("Invalid token expiry")
		return nil,false
	}

	// Extract specific claims
	_claims := map[string]interface{}{
		"username": claims["username"].(string),
		"email": claims["email"].(string),
		"expire_at" : expire_at,
	}

	expiredAt := time.Unix(int64(expire_at), 0)

	// Check if the token has expired
	if expiredAt.Before(time.Now()) {
		fmt.Println("Token has expired!")
		return _claims,false
	} 
	return _claims,true
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