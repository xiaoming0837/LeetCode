package main

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {

}

func reverse(head *ListNode) *ListNode {
	newHead := head.Next
	head.Next = reverse(newHead.Next)
	newHead.Next = head
	return newHead
}
