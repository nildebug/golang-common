package tools

import (
	"math/rand"
	"time"
)

// RandomInt
//
//	@Description: 生成随机数
//	@param min
//	@param max
//	@return int
func RandomInt(min, max int) int {
	// 设置随机数种子
	rand.Seed(time.Now().UnixNano())

	// 计算随机数的范围
	rangeSize := max - min + 1

	// 生成随机数
	randomInt := rand.Intn(rangeSize) + min

	return randomInt
}
