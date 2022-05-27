package leetcode

// 给你一个链表，两两交换其中相邻的节点，并返回交换后链表的头节点。你必须在不修改节点内部的值的情况下完成本题（即，只能进行节点交换）。
func swapPairs(head *ListNode) *ListNode {
	// 思路
	// 把每次两个节点交换的数据看成是一个整体，只对这两个节点进行数据交换, 这样来看的话， 就能把一个链的问题分解成2个节点的问题
	// 然后后续节点只需要回调调用即可解决。
	// 如果这个节点没有下一个节点， 表明是单独的一个节点，直接返回
	//
	if head == nil || head.Next == nil {
		return head
	}

	next := head.Next
	head.Next = swapPairs(next.Next)
	next.Next = head

	return next
}
