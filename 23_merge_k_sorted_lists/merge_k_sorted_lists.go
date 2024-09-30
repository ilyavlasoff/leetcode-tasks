package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// You are given an array of k linked-lists lists,
// each linked-list is sorted in ascending order.
//
// Merge all the linked-lists into one sorted
// linked-list and return it.
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	for len(lists) > 1 {
		var iters int
		for i := 0; i < len(lists); i = i + 2 {
			if i == len(lists)-1 {
				lists[iters] = lists[i]
				iters++
				continue
			}
			res := mergeList(lists[i], lists[i+1])
			lists[iters] = res
			iters++
		}
		lists = lists[:iters]
	}

	return lists[0]
}

func mergeList(a *ListNode, b *ListNode) *ListNode {
	if a == nil {
		return b
	}
	if b == nil {
		return a
	}

	a = &ListNode{Val: min(a.Val, b.Val), Next: a}
	aHead := a

	for a != nil && b != nil {
		if a.Next == nil {
			a.Next = b
			break
		}

		if b == nil {
			break
		}

		if a.Val <= b.Val && a.Next.Val > b.Val {
			bCur := b
			b = b.Next

			bCur.Next = a.Next
			a.Next = bCur
		}

		a = a.Next
	}
	return aHead.Next
}
