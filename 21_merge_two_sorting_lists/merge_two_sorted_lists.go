package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	n := &ListNode{}
	root := n

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			n.Next = list1
			list1 = list1.Next
		} else {
			n.Next = list2
			list2 = list2.Next
		}

		n = n.Next
	}

	if list1 == nil && list2 != nil {
		n.Next = list2
	}
	if list2 == nil && list1 != nil {
		n.Next = list1
	}

	return root.Next
}
