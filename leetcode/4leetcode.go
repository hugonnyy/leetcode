package leetcode

//给定两个大小分别为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。请你找出并返回这两个正序数组的 中位数 。
//算法的时间复杂度应该为 O(log (m+n)) 。

// *****
// 要解决此题，需要理解什么是中位数
// 如果有7个数，需要排除前面3个数，取第4个数
// 如果有9个数，需要排除前面4个数，取第5个数
// 如果有8个数，需要排除前面3个数，取4，5个数  ===> 转换成 分别获取第4位数 和 获取 第5位数
// 如果有10个数，需要排除前面4个数字，取5，6个数  ===> 转换成 分别获取第5位数 和 获取 第6位数
//
// 对于取两个数组的中位数，不能直接排除，需要则中法排除： 原先是排除4个，分别分到两个数组中，就是一个数组排除2个
//
// *****
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	totalLength := len(nums1) + len(nums2)

	if totalLength%2 == 1 { // 有中位数
		midIndex := totalLength/2 + 1
		return float64(getKthElement(nums1, nums2, midIndex))
	} else {
		midIndex := totalLength / 2
		return float64(getKthElement(nums1, nums2, midIndex)+getKthElement(nums1, nums2, midIndex+1)) / 2.0
	}
}

func getKthElement(nums1, nums2 []int, k int) int {
	index1, index2 := 0, 0

	for {
		if index1 == len(nums1) {
			return nums2[index2+k-1]
		}
		if index2 == len(nums2) {
			return nums1[index1+k-1]
		}

		if k == 1 {
			return min(nums1[index1], nums2[index2])
		}

		half := k / 2

		newIndex1 := min(index1+half, len(nums1)) - 1
		newIndex2 := min(index2+half, len(nums2)) - 1

		pivot1, pivot2 := nums1[newIndex1], nums2[newIndex2]

		if pivot1 <= pivot2 {
			k -= newIndex1 - index1 + 1
			index1 = newIndex1 + 1
		} else {
			k -= newIndex2 - index2 + 1
			index2 = newIndex2 + 1
		}

	}

	return 0
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
