package main

func main() {

}

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	ans := head
	for head.Next != nil && head.Next.Next != nil {
		n1 := head
		n2 := head.Next
		n1.Next = n2.Next
		n2.Next = head
		head = n2.Next
	}
	return ans
}
