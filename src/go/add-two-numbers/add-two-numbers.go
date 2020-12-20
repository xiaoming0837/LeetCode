package main

import "fmt"

func main() {
	fmt.Fprintf("%d", 23/10)
}

/**
 * Defintion for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

// add two numbers
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	p := dummy
	tmp := 0
	for l1 != nil || l2 != nil || tmp != 0 {
		p.Val += tmp
		tmp = 0
		if l1 != nil {
			p.Val += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			p.Val += l2.Val
			l2 = l2.Next
		}
		tmp = p.Val / 10
		p.Val = p.Val % 10
		if l1 != nil || l2 != nil || tmp != 0 {
			p.Next = &ListNode{}
			p = p.Next
		}
	}
	return dummy
}
