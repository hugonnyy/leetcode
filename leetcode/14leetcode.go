package leetcode

// 编写一个函数来查找字符串数组中的最长公共前缀。
//
//如果不存在公共前缀，返回空字符串 ""。
// 此题运行速度3次都超过100%的人
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	strLength := len(strs)

	index := -1

	for i := 0; i < len(strs[0]); i++ {
		word := strs[0][i]
		for j := 1; j < strLength; j++ {
			if i >= len(strs[j]) {
				return strs[0][:index+1]
			}

			if word != strs[j][i] {
				return strs[0][:index+1]
			}
		}

		index = i
	}

	return strs[0][:index+1]
}
