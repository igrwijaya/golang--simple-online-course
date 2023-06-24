package utils

import (
	"github.com/golang-jwt/jwt/v5"
	"os"
)

type AppClaims struct {
	Id      int
	Name    string
	Email   string
	IsAdmin bool
	jwt.RegisteredClaims
}

func CreateJwtToken(claims AppClaims) string {
	jwtKey := os.Getenv("JWT_SECRET")
	jwtKeyByte := []byte(jwtKey)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, errOnSigningToken := token.SignedString(jwtKeyByte)

	if errOnSigningToken != nil {
		panic("Cannot create JWT token. " + errOnSigningToken.Error())
	}

	return tokenString
}
