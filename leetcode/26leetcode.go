package leetcode

// 给你一个 升序排列 的数组 nums ，请你 原地 删除重复出现的元素，使每个元素 只出现一次 ，返回删除后数组的新长度。元素的 相对顺序 应该保持 一致 。
//
//由于在某些语言中不能改变数组的长度，所以必须将结果放在数组nums的第一部分。更规范地说，如果在删除重复项之后有 k 个元素，那么 nums 的前 k 个元素应该保存最终结果。
//
//将最终结果插入 nums 的前 k 个位置后返回 k 。
//
//不要使用额外的空间，你必须在 原地 修改输入数组 并在使用 O(1) 额外空间的条件下完成。
// [1,1,2]
// -----------
// [1,1]
// -----------
// [1,2]

// [0,0,1,1,1,2,2,3,3,4]
// ---------------------
// [0,1,1,2,3,4]
// ---------------------
// [0,1,2,3,4]
// 这种算法效率太低了，
//func removeDuplicates(nums []int) int {
//	if len(nums) == 0 || len(nums) == 1 {
//		return len(nums)
//	}
//
//	length := len(nums)
//
//	for i := 1; i < length; i++ {
//		// 当前元素跟上一个元素比是否一致
//		if nums[i - 1] != nums[i] {
//			continue
//		}
//
//		// 当前元素与上一个元素相等，把后续元素都往前移
//		for j := i; j < length - 1; j++ {
//			nums[j] = nums[j + 1]
//		}
//		length --
//		i = i -1
//	}
//
//	return length
//}

// 查看参考答案， 可以使用双指针方法
func removeDuplicates(nums []int) int {
	left := 0

	for right := 1; right < len(nums); right++ {
		if nums[left] != nums[right] {
			nums[left+1] = nums[right]
			left++
		}
	}

	return left + 1
}
