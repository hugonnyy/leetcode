package leetcode

// 给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。
//
//请必须使用时间复杂度为 O(log n) 的算法。
func searchInsert(nums []int, target int) int {
	// 思路
	// 已知条件： 数组已排序， 需要查找某一个数字
	// 采用二分法查找
	left := 0
	right := len(nums) - 1
	ans := len(nums)

	for left <= right {
		mid := (right-left)>>1 + left
		if target <= nums[mid] {
			ans = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return ans
}
