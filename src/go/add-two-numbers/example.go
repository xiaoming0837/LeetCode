package main

/**
 * Defintion for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carryFlag := 0
	var rs *ListNode
	var cursor *ListNode
	for l1 != nil || l2 != nil || carryFlag > 0 {
		firstV := 0
		secondV := 0
		if l1 != nil {
			firstV = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			secondV = l2.Val
			l2 = l2.Next
		}
		num := firstV + secondV + carryFlag
		carryFlag = num / 10
		num = num % 10
		currentN := ListNode{num, nil}
		if rs == nil {
			rs = &currentN
			cursor = rs
		} else {
			cursor.Next = &currentN
			cursor = cursor.Next
		}
	}
	return rs
}
