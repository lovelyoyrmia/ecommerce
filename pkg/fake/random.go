package fake

import (
	"fmt"
	"math/rand"
)

func RandomString() string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	s := make([]rune, 10)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}

func RandomInt(min, max int) int {
	return rand.Intn(max-min) + min
}

func RandomEmail() string {
	username := RandomString()
	domain := "test.com"
	email := fmt.Sprintf("%s@%s", username, domain)
	return email
}

func RandomStock() int {
	return RandomInt(0, 30)
}

func RandomPrice() int {
	return RandomInt(0, 20000)
}
