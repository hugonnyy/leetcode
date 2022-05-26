package leetcode

import (
	"math"
	"strconv"
)

// 给你一个 32 位的有符号整数 x ，返回将 x 中的数字部分反转后的结果。
//
//如果反转后整数超过 32 位的有符号整数的范围 [−231,  231 − 1] ，就返回 0。
//
//假设环境不允许存储 64 位整数（有符号或无符号）。
func reverse(x int) int {
	symbol := 1
	if x < 0 {
		symbol = -1
		x *= -1
	}

	xBytes := []byte(strconv.Itoa(x))
	for i := 0; i < len(xBytes)/2; i++ {
		tmp := xBytes[i]
		xBytes[i] = xBytes[len(xBytes)-1-i]

		xBytes[len(xBytes)-1-i] = tmp
	}

	a, err := strconv.Atoi(string(xBytes))
	if err != nil {
		return 0
	}

	x = a * symbol

	if x < math.MinInt32 || x > math.MaxInt32 {
		return 0
	}

	return x
}
