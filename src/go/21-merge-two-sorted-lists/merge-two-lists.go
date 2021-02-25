package main

func main() {
	fmt.Println(mergeTwoLists([]))
}

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil{
		return l2
	} else if l2 == nil {
		return l1
	}
	ans := new(ListNode)
	dummy := ans
	for l1 != nil || l2 != nil {
		if l1==nil {
			dummy.Next = &ListNode{l2.Val, nil}
			l2 = l2.Next
		} else if l2 == nil {
			dummy.Next = &ListNode{l1.Val, nil}
			l1 = l1.Next
		} else if l1.Val < l2.Val {
			dummy.Next = &ListNode{l1.Val, nil}
			l1 = l1.Next
		} else {
			dummy.Next = &ListNode{l2.Val, nil}
			l2 = l2.Next
		}
		dummy = dummy.Next
	}
	return ans.Next
}
