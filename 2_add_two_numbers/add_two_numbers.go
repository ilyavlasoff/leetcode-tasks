package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// You are given two non-empty linked lists representing two non-negative integers.
// The digits are stored in reverse order, and each of their nodes contains a single digit.
// Add the two numbers and return the sum as a linked list.
//
// You may assume the two numbers do not contain any leading zero, except the number 0 itself.
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var carry int
	var head ListNode
	var cur = &head

	for l1 != nil || l2 != nil {
		val := carry
		if l1 != nil {
			val += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			val += l2.Val
			l2 = l2.Next
		}

		carry = 0
		if val/10 > 0 {
			carry = val / 10
			val = val % 10
		}

		node := &ListNode{Val: val}
		cur.Next = node
		cur = node
	}

	if carry > 0 {
		cur.Next = &ListNode{Val: carry}
	}

	return head.Next
}
