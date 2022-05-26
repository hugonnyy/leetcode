package leetcode

import "strconv"

// 给你一个整数 x ，如果 x 是一个回文整数，返回 true ；否则，返回 false 。
//
//回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。
//
//例如，121 是回文，而 123 不是。
// 效率和内存都不行
// leetcode 上的效率有问题， 同一份代码， 点击就会有速度快和速度慢的区别
func isPalindrome(x int) bool {
	xStr := strconv.Itoa(x)
	length := len(xStr)
	for i := 0; i < length; i++ {
		if xStr[i] != xStr[length-1-i] {
			return false
		}
	}

	return true
}
