package leetcode

// 给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度。
// 思路， 采用双指针方法，一个指针指向头，一个指针指向尾
func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	// 想要速度快就要用内存搜索
	sBytes := []byte(s)

	cache := make(map[byte]int) // key: byte, value: index
	leftIndex := 0
	maxLength := 0

	for i := 0; i < len(s); i++ {
		index, exist := cache[sBytes[i]]
		if exist { // 缓存中不存在该字符，
			cacheIndex := index + 1
			leftIndex = max(cacheIndex, leftIndex)
		}

		cache[sBytes[i]] = i
		maxLength = max(maxLength, i-leftIndex+1)
	}

	return maxLength
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
