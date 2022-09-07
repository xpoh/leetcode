package main

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
	result := &ListNode{}
	p1 := l1
	p2 := l2
	p3 := result
	for p1 != nil || p2 != nil {
		if p1 != nil {
			p3.Val = p3.Val + p1.Val
			p1 = p1.Next
		}
		if p2 != nil {
			p3.Val = p3.Val + p2.Val
			p2 = p2.Next
		}
		if p1 != nil || p2 != nil {
			p3.Next = &ListNode{}
		}
		if p3.Val > 9 {
			if p3.Next == nil {
				p3.Next = &ListNode{}
			}
			p3.Next.Val = 1
			p3.Val = p3.Val - 10
		}
		p3 = p3.Next
	}
	return result
}
