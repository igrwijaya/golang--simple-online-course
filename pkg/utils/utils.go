package utils

import "math/rand"

func RandString(length int) string {
	var letterRune = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ123456789")

	char := make([]rune, length)

	for i := range char {
		char[i] = letterRune[rand.Intn(len(letterRune))]
	}

	return string(char)
}
