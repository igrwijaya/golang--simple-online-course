package jwt_utils

import (
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"os"
)

type AppClaims struct {
	Id      uint
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

func HasValidPassword(hasPassword string, plainPassword string) bool {
	errHashCompare := bcrypt.CompareHashAndPassword([]byte(hasPassword), []byte(plainPassword))

	if errHashCompare != nil {
		return false
	}

	return true
}
