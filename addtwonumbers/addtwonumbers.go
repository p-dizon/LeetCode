package addtwonumbers

import "github.com/p-dizon/LeetCode/linkedlist"

func addTwoNumbers(l1 *linkedlist.ListNode, l2 *linkedlist.ListNode) *linkedlist.ListNode {
	var sum, node *linkedlist.ListNode
	var carry int
	for p1, p2 := l1, l2; p1 != nil || p2 != nil || carry != 0; {
		if node == nil {
			node = new(linkedlist.ListNode)
			sum = node
		} else {
			node.Next = new(linkedlist.ListNode)
			node = node.Next
		}

		var a, b int
		if p1 != nil {
			a = p1.Val
			p1 = p1.Next
		}
		if p2 != nil {
			b = p2.Val
			p2 = p2.Next
		}

		s := carry + a + b
		carry = s / 10
		node.Val += s % 10
	}
	return sum
}
