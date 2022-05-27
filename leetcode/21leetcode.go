package leetcode

// 将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	result := &ListNode{}
	cur := result

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			cur.Next = list1
			cur = list1
			list1 = list1.Next
		} else {
			cur.Next = list2
			cur = list2
			list2 = list2.Next
		}
	}

	if list1 != nil {
		cur.Next = list1
	}

	if list2 != nil {
		cur.Next = list2
	}

	return result.Next
}
