package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	var res bool

	if head == nil {
		return false
	}

	seen := make(map[*ListNode]bool)

	for {
		_, ok := seen[head]
		if ok {
			return true
		}
		seen[head] = true

		if head.Next == nil {
			break
		}

		head = head.Next
	}
	// TODO: 走訪節點，用 map 記錄到過的節點本身。
	// 想想：何時回傳 true？何時能確定回傳 false？
	return res
}

func hasCycleFastSlow(head *ListNode) bool {
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

	// TODO: 讓慢指標每輪走一步，快指標每輪走兩步。
	// 每輪移動前，先確認快指標還能安全地往前走。
	return res
}

func main() {
	a := &ListNode{Val: 1}
	b := &ListNode{Val: 2}
	c := &ListNode{Val: 1} // 值與 a 相同，但不是同一個節點。
	a.Next, b.Next, c.Next = b, c, b
	fmt.Println(hasCycle(a))         // 預期：true
	fmt.Println(hasCycleFastSlow(a)) // 預期：true

	c.Next = nil
	fmt.Println(hasCycle(a))           // 預期：false
	fmt.Println(hasCycleFastSlow(a))   // 預期：false
	fmt.Println(hasCycle(nil))         // 預期：false
	fmt.Println(hasCycleFastSlow(nil)) // 預期：false
}
