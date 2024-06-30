package auth

import (
	"fmt"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
)

// TODO: implement jwt.go

// use jwt for authentication

type UserClaims struct {
	UserId string `json:"userId"`
	jwt.RegisteredClaims
}

func ParseToken(token string, secret string) (string, error) {

	return "", nil
}

func CreateToken(userId string, secret string) (string, error) {
	now := jwt.NewNumericDate(time.Now())
	expires := jwt.NewNumericDate(time.Now().Add(24 * time.Hour))
	claims := UserClaims{
		userId,
		jwt.RegisteredClaims{
			ExpiresAt: expires,
			IssuedAt:  now,
			NotBefore: now,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", fmt.Errorf("Error signing token: %w", err)
	}

	return signedToken, nil
}

func createTokenWithClaims(userId string, secret string, userClaims UserClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, userClaims)
	signedToken, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", fmt.Errorf("Error signing token: %w", err)
	}

	return signedToken, nil
}
