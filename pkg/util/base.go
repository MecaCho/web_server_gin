package util

import (
	"math/rand"
	"time"
)

// GenPwd generate a random password
func GenPwd(num int) string {
	rand.NewSource(time.Now().UnixNano())
	rand.Seed(time.Now().UnixNano())
	res := make([]byte, num)
	seedStr := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	for i := 0; i < num; i++ {
		seq := rand.Intn(len(seedStr))
		res[i] = seedStr[seq]
	}
	return string(res)
}
