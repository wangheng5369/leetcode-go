package linkedlist

// p141 环形链表
// 给你一个链表的头节点 head，判断链表中是否有环。
// 如果链表中存在环，则返回 true，否则返回 false。
//
// 提示：
// - 链表中结点的数目范围是 [0, 10^5]
// - -10^5 <= Node.val <= 10^5

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	slow, fast := head, head
	for fast != nil && slow != nil {
		if fast.Next == nil || fast.Next.Next == nil {
			return false
		}
		fast = fast.Next.Next
		slow = slow.Next
	}
	return true
}

func hasCycle1(head *ListNode) bool {
	m := make(map[*ListNode]struct{})
	curr := head
	for curr != nil {
		if _, ok := m[curr]; ok {
			return true
		} else {
			m[curr] = struct{}{}
			curr = curr.Next
		}
	}
	return false
}
