# 0206. Reverse Linked List - 原始作答

## 2026-10-12 作答；2026-10-13 歸檔

### 完整原始 `main.go`

```go
package main

import "fmt"

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

func main() {
	head := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}}
	// 預期：3 2 1
	for node := reverseList(head); node != nil; node = node.Next {
		fmt.Print(node.Val, " ")
	}
	fmt.Println()

	// 預期：true
	fmt.Println(reverseList(nil) == nil)
}
```

### 本輪思路與修正

- 最初提出將節點記入陣列後反向修改連結，時間 O(n)、額外空間 O(n)，並指出原 head 變成尾端後 `Next` 應為 nil。
- 經引導改採三指標原地反轉；使用者正確指出修改 `Next` 前必須保存原下一節點，以免失去後續路徑。
- 初稿對 nil head 解參考、單節點回傳 nil，且提前停止；第二稿仍提前停止。第三稿自行修正，已執行的多節點、單節點、兩節點與 nil 案例均通過。
- 使用者口述：`curr` 到 nil 時 `prev` 指向反轉後新 head；時間 O(n)、額外空間 O(1)。nil／單節點雖通過程式驗證，尚未由使用者獨立口述邊界理由。
