# 0206. Reverse Linked List

## 題目

輸入單向鏈結串列的 `head`，反轉所有節點的 `Next` 方向並回傳新頭；空串列回傳 nil。Easy。

## 本輪最終解法

```go
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

```

最終解法與使用者作答的型別及函式相同，只移除範例輸出的 `main`。完整原始作答見 [my_attempt.md](my_attempt.md)。

## 推導與正確性

先走訪並把 `*ListNode` 存到陣列，再依反向順序重接，可用 O(n) 時間／O(n) 空間。原地做法只保留 `prev`（已反轉部分的新頭）、`curr`（待處理節點）、`next`（尚未處理的後續節點）。改 `curr.Next` 前先保存後續連結，每處理一節點就把它接到已反轉部分前端。已處理部分始終是原前綴的反轉，後續節點仍可由 `next` 到達；全部處理完 `curr` 為 nil，`prev` 指向原尾端，即新 head。此完整不變量由筆記整理；使用者已口述保存後續節點與 `prev` 成為新頭的理由。

## 複雜度

令 n 為節點數：時間 O(n)，每節點處理一次；額外空間 O(1)，僅固定數量的指標，沒有遞迴堆疊。不計輸入串列本身。

## 驗證與邊界

| 輸入 | 預期 | 實際 | 驗證 |
|---|---|---|---|
| `1 -> 2 -> 3 -> nil` | `3 -> 2 -> 1 -> nil` | `3 2 1` | `timeout 5s go run main.go` |
| `nil` | `nil` | `true`（回傳 nil） | `go run main.go` 與臨時 Go 程式 |
| `9 -> nil` | 同一節點，`Next=nil` | `true 9 true` | 臨時 Go 程式 |
| `1 -> 2 -> nil` | `2 -> 1 -> nil` | `2 1 true` | 臨時 Go 程式 |

第一稿對三節點輸出 `2 3` 後在 nil 輸入 panic；第二稿三節點只輸出 `1`。上表為第三稿修正後的實際結果，案例通過不等於窮舉證明。

## 面試口述版

用 `prev` 保存已反轉部分、`curr` 走訪原串列。每輪先記住原下一節點，將目前節點指回 `prev`，再推進兩個指標。`curr` 到 nil 時 `prev` 是新頭。時間 O(n)，額外空間 O(1)。

## 待複習

- nil 與單節點的處理雖有執行證據，尚未由使用者獨立口述原因。
- 可練習更簡潔的 `for curr != nil` 迴圈寫法；本輪保留已驗證的使用者作答。
