package linkedlist

import (
	"fmt"
)

// ListNode is a node in a linked list
type ListNode struct {
	Val  int
	Next *ListNode
}

// ConvertToList converts an int slice into a ListNode
func ConvertToList(vals []int) *ListNode {
	var list, node *ListNode
	for _, val := range vals {
		if node == nil {
			node = new(ListNode)
			list = node
		} else {
			node.Next = new(ListNode)
			node = node.Next
		}
		node.Val = val
	}

	return list
}

// IsEqual returns true if the linked lists contain the same values
// false otherwise
func IsEqual(l1, l2 *ListNode) (bool, error) {
	for i := 0; l1 != nil && l2 != nil; i++ {
		if l2.Val != l1.Val {
			return false, fmt.Errorf("Different values for index %v; l1.Val = [%v], l2.Val = [%v]", i, l1.Val, l2.Val)
		}
		l2 = l2.Next
		l1 = l1.Next
	}
	if l1 != l2 {
		return false, fmt.Errorf("Node count mismatch")
	}
	return true, nil
}
