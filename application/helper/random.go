package helper

import (
	"math/rand"
	"time"
)

func GenerateToken(n int) string {
	rand.Seed(time.Now().UnixNano())

	var letterRunes = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
