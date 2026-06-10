package linkedlist

// p206 反转链表
// 给你单链表的头节点 head，请你反转链表，并返回反转后的链表。
//
// 示例 1：
// 输入：head = [1, 2, 3, 4, 5]
// 输出：[5, 4, 3, 2, 1]
//
// 提示：
// - 链表中节点数目范围是 [0, 5000]
// - -5000 <= Node.val <= 5000

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head
	for curr != nil {
		t := curr.Next
		curr.Next = prev
		prev = curr
		curr = t
	}
	return prev
}
