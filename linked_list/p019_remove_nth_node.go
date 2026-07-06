package linkedlist

// p019 删除链表的倒数第 N 个结点
// 给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。
//
// 示例 1：
// 输入：head = [1, 2, 3, 4, 5], n = 2
// 输出：[1, 2, 3, 5]
//
// 提示：
// - 链表中结点的数目为 sz
// - 1 <= sz <= 30
// - 1 <= n <= sz

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}
	p, q := dummy, dummy
	for i := 0; i <= n; i++ {
		p = p.Next
	}
	for p != nil {
		p = p.Next
		q = q.Next
	}
	return dummy.Next
}

func removeNthFromEnd1(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}
	nodes := []*ListNode{}
	for node := dummy; node != nil; node = node.Next {
		nodes = append(nodes, node)
	}
	prev := nodes[len(nodes)-1-n]
	prev.Next = prev.Next.Next
	return dummy.Next
}
