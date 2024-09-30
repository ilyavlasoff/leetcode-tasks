package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}

	var (
		delPred = head
		s       int
	)

	for cur := head; cur != nil; cur = cur.Next {
		if s > n {
			delPred = delPred.Next
		}

		s++
	}

	if s < n {
		return head
	}
	if s == n {
		return head.Next
	}

	if delPred.Next == nil {
		return nil
	}

	delPred.Next = delPred.Next.Next

	return head
}
