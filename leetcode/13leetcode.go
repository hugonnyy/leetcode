package leetcode

// 罗马数字包含以下七种字符: I， V， X， L，C，D 和 M。
// 给定一个罗马数字，将其转换成整数。
func romanToInt(s string) int {
	result := 0

	lastV := 0
	for i := len(s) - 1; i >= 0; i-- {
		curV := cache[s[i]]
		if curV < lastV {
			result -= curV
		} else {
			result += curV
		}

		lastV = curV
	}

	return result
}

var cache = map[byte]int{
	'M': 1000,
	'D': 500,
	'C': 100,
	'L': 50,
	'X': 10,
	//cache["IX"] = 9
	'V': 5,
	//cache["IV"] = 4
	'I': 1,
}
