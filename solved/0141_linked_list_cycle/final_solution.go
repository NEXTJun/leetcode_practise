package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	var res bool

	if head == nil {
		return false
	}

	h1 := head
	h2 := head

	for {
		if h1.Next == nil {
			break
		}
		if h2.Next == nil || h2.Next.Next == nil {
			break
		}
		h1 = h1.Next
		h2 = h2.Next.Next
		if h1 == h2 {
			return true
		}
	}

	return res
}
