package util

import (
	"math/rand"
	"time"
)

func RandString(length int) string {
	rand.Seed(time.Now().Unix())
	chars := "abcdefghijklmnopqrstuvwxyz0123456789"
	nonce := make([]byte, length)
	for i := 0; i < length; i++ {
		nonce[i] = byte(chars[rand.Intn(len(chars))])
	}
	return string(nonce)
}
