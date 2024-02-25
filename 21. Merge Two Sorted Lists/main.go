package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(a *ListNode, b *ListNode) *ListNode {
	root := &ListNode{}
	cur := root

	for a != nil || b != nil {
		if (a != nil && b != nil && a.Val <= b.Val) || b == nil {
			cur.Next = a
			a = a.Next
		} else if b != nil {
			cur.Next = b
			b = b.Next
		}
		cur = cur.Next
	}
	return root.Next
}
