package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	var prev, curr, next *ListNode
	next = head.Next
	curr = head
	prev = nil

	for {
		curr.Next = prev
		prev = curr
		curr = next
		if curr == nil {
			break
		}
		next = next.Next
	}

	return prev
}

