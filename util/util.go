package util

import (
	"math/rand"
	"time"
)

//RandomString 随机生成用户名
func RandomString(n int) string {
	var letters = []byte("qwertyuiopasdfghjklzxcvbnm")
	result := make([]byte, n)

	rand.Seed(time.Now().Unix())
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}
