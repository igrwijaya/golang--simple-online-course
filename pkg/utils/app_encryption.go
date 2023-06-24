package utils

import "golang.org/x/crypto/bcrypt"

func EncryptPassword(password string) string {
	hashedPassword, errHashedPassword := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost)

	if errHashedPassword != nil {
		panic("Cannot encrypt the Password")
	}

	return string(hashedPassword)
}
