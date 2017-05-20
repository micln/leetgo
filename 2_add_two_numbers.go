package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	tail := head

	v1 := 0
	v2 := 0
	g := 0
	sum := 0

	for {
		l1, v1 = current(l1)
		l2, v2 = current(l2)

		sum = (v1 + v2 + g) % 10
		g = (v1 + v2 + g) / 10

		tail = append(tail, sum)

		if l1 == nil && l2 == nil && g == 0 {
			break
		}
	}

	return head.Next
}

func append(tail *ListNode, v int) *ListNode {
	n := &ListNode{Val: v}
	tail.Next = n

	return n
}

func current(tail *ListNode) (*ListNode, int) {
	if tail == nil {
		return nil, 0
	}

	return tail.Next, tail.Val
}
