package leetcode

// 给定一个长度为 n 的整数数组 height 。有 n 条垂线，第 i 条线的两个端点是 (i, 0) 和 (i, height[i]) 。
//
//找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。
//
//返回容器可以储存的最大水量。
//
//说明：你不能倾斜容器。
func maxArea(height []int) int {
	// 思路：
	// 首先要确认， 肯定要遍历所有的节点， 因为可能有的节点上的数据是无穷大的,这样就会使得结果是最大的。

	start := 0
	end := len(height)
	area := 0

	for start < end {
		w := end - start - 1
		h := min(height[start], height[end])

		area = max(area, w*h)

		if height[start] < height[end] {
			start++
		} else {
			end--
		}

	}

	return area
}

//func min(a, b int) int {
//	if a < b {
//		return a
//	}
//
//	return b
//}
//
//func max(a, b int) int {
//	if a > b {
//		return a
//	}
//
//	return b
//}
