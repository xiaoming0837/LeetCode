package main

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	var pre *ListNode
	dummy := &ListNode{0,head}
	dummy.Next = next(pre,head, k)
	return dummy
}

func next(pre,head *ListNode, k int) *ListNode {
	if head == nil {
		return pre
	}
	pre,head = reverse(pre,head,k)
	return next(pre,head, k)
}

func reverse(head *ListNode, k int) (*ListNode ,*ListNode) {
	curr := head
  var pre *ListNode
	for ; k > 0 && curr != nil; k-- {
		next := curr.Next
		curr.Next = pre
		pre = curr
		curr = next
	}
	return pre,curr
}
