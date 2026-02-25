package linkedlist

// You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order, and each of their nodes contains a single digit. Add the two numbers and return the sum as a linked list.

// You may assume the two numbers do not contain any leading zero, except the number 0 itself.

// Input: l1 = [2,4,3], l2 = [5,6,4]
// Output: [7,0,8]
// Explanation: 342 + 465 = 807.

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := 0
	head := &ListNode{Val: 0, Next: nil}
	current := head
	reversedL1 := l1
	reversedL2 := l2
	for reversedL1 != nil || reversedL2 != nil {
		val1 := 0
		val2 := 0
		if reversedL1 != nil {
			val1 = reversedL1.Val
			reversedL1 = reversedL1.Next
		}
		if reversedL2 != nil {
			val2 = reversedL2.Val
			reversedL2 = reversedL2.Next
		}
		sum := val1 + val2 + carry
		carry = sum / 10
		sum = sum % 10
		newNode := &ListNode{Val: sum, Next: nil}
		current.Next = newNode
		current = newNode
	}
	if carry > 0 {
		current.Next = &ListNode{Val: carry, Next: nil}
	}
	return head.Next
}
