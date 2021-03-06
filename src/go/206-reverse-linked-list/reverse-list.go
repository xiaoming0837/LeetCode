package main

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	dummy := &ListNode{0, head}
	for dummy.Next != nil && dummy.Next.Next != nil {
		n1 := dummy.Next
		n2 := dummy.Next.Next
		n2.Next = n1
		n1 = nil
		dummy = n2
	}
	return dummy.Next
}
