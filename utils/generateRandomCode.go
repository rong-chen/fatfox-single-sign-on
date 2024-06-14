package utils

import (
	"math/rand"
	"time"
)

func GenerateRandomCode(length int) string {
	// 生成指定长度的随机数验证码
	// 定义验证码字符集
	const charset = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	// 初始化随机数种子
	rand.Seed(time.Now().UnixNano())
	// 生成验证码
	code := make([]byte, length)
	for i := range code {
		code[i] = charset[rand.Intn(len(charset))]
	}
	return string(code)
}
