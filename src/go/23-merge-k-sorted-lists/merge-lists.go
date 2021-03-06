package main

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	return mergeSort(lists, 0, len(lists)-1)
}

func mergeSort(lists []*ListNode, left, right int) *ListNode {
	if left == right {
		return lists[right]
	}
	if left > right {
		return nil
	}
	mid := (left + right) >> 1
	return mergeTwoLists(mergeSort(lists, left, mid), mergeSort(lists, mid+1, right))
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	}
	ans := new(ListNode)
	dummy := ans
	for l1 != nil || l2 != nil {
		if l1 == nil {
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
