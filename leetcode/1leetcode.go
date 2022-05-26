package leetcode

// 给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。
//
//你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。
//
//你可以按任意顺序返回答案。
func twoSum(nums []int, target int) []int {

	// 为了达到速度快的目的， 需要用空间换取时间，采用数据结构可以做到这一点
	cache := make(map[int]int)

	// 判断当前缺少的数值是否在内存中，
	// 如果存在，直接取出来
	for i := 0; i < len(nums); i++ {
		index, exist := cache[target-nums[i]]
		if exist {
			return []int{i, index}
		}

		cache[nums[i]] = i
	}

	return []int{}
}
