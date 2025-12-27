package utils

import (
	"math/rand"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// GenerateRandomPassword returns a random password with length in the
// inclusive range [minLen, maxLen] using a small charset. It seeds the
// global random generator with the current time and returns the generated
// password as a string. This helper is intended for example/demo use.
func GenerateRandomPassword(minLen, maxLen int) string {
	rand.Seed(time.Now().UnixNano())

	length := rand.Intn(maxLen-minLen+1) + minLen
	password := make([]byte, length)

	for i := range password {
		password[i] = charset[rand.Intn(len(charset))]
	}

	Debugf("Generated random password: %s", string(password))
	return string(password)
}
