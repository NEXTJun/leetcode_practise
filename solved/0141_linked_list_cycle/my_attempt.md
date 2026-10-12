# 0141. Linked List Cycle - 原始作答

## 2026-10-13 map 版本（加入優化框架前）

```go
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

func main() {
	a := &ListNode{Val: 1}
	b := &ListNode{Val: 2}
	c := &ListNode{Val: 1} // 值與 a 相同，但不是同一個節點。
	a.Next, b.Next, c.Next = b, c, b
	fmt.Println(hasCycle(a)) // 預期：true

	c.Next = nil
	fmt.Println(hasCycle(a))   // 預期：false
	fmt.Println(hasCycle(nil)) // 預期：false
}
```

- 使用者口述：節點值可重複，應以節點本身作為 map 的鍵。
- 執行結果：原範例有環／無環／nil 輸出 true、false、false；補充單節點無環、自環、同值不同節點無環輸出 false、true、false。
- 此版本時間 O(n)、額外空間 O(n)；已保存，接續嘗試快慢指標 O(1) 空間版。

## 2026-10-13 快慢指標版本（完整最終 `main.go`）

```go
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
```

- 使用者自行加入快慢指標版；初稿因未確認 `h2.Next` 就讀 `h2.Next.Next`，無環三節點輸入執行時 panic。
- 修正後，原範例及單節點無環、自環、兩節點成環、同值不同節點無環、nil 案例均通過。
- 使用者口述：移動後相遇回 true；先到尾端回 false；時間 O(n)、額外空間 O(1)。自環已測試，但使用者未另作單獨口述。
