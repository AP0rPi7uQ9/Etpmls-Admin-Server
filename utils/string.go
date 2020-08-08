package utils

import (
	"math/rand"
	"time"
)

// Generate random strings
// 生成随机字符串
func GenerateRandomString(l int) string {
	var code = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/~!@#$%^&*()_="

	data := make([]byte, l)
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < l; i++ {
		idx := rand.Intn(len(code))
		data[i] = code[idx]
	}
	return string(data)
}