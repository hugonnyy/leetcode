package leetcode

// 罗马数字包含以下七种字符： I， V， X， L，C，D 和 M。
// 给你一个整数，将其转为罗马数字。
func intToRoman(num int) string {
	// 思路：
	// 1.建立数字与字符的联系
	// 2.特殊数字采用也是数字与字符的联系
	// 3.从大到小循环遍历列举的所有数据
	result := ""
	for _, l := range list {
		if num > l.value {
			num -= l.value
			result += l.symbol
		}

		if num == 0 {
			break
		}
	}

	return result
}

var list = []struct {
	value  int
	symbol string
}{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}
