# 0141. Linked List Cycle

## 題目

給定單向鏈結串列 `head`，若沿著 `Next` 會重新到達同一節點則回傳 `true`，若走到 nil 則回傳 `false`。節點的 `Val` 可以重複；判定依據是節點身分。Easy。

## 最終 Go 解法

```go
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
```

此檔沿用使用者快慢指標實作，只把函式名改成題目使用的 `hasCycle` 並移除 scaffold TODO 註解。完整 map 版和快慢指標版原始 `main.go` 均在 [my_attempt.md](my_attempt.md)。

## 推導與正確性

直觀法用 `map[*ListNode]bool` 記已訪問的節點；重訪同一指標即有環，走到 nil 即無環，時間 O(n)、額外空間 O(n)。快慢指標法讓慢指標每輪走一步，快指標每輪走兩步。若無環，快指標最終無法再走並回 false。若有環，設入環前有 a 個節點、環長 c；慢指標在 a 輪後進環，此時快指標也在環內。兩者相對位置每輪改變一格，在環內至多再 O(c) 輪會相遇並回 true。總輪數 O(a+c)=O(n)。此一般性上界由教學整理；使用者能說出會追上、四節點環的精確最多三輪，但未獨立完整口述 a+c 證明。

## 複雜度

令 n 為從 head 可到達的不同節點數。map 版時間 O(n)、額外空間 O(n)；快慢指標版時間 O(n)、額外空間 O(1)。兩者皆無遞迴堆疊；輸入串列本身不計入額外空間。

## 驗證與邊界

| 輸入情況 | 預期 | 實際 | 驗證 |
|---|---|---|---|
| `a -> b -> c -> b` | true | map true、快慢 true | `timeout 5s go run main.go` |
| `a -> b -> c -> nil`（a 與 c 同值） | false | map false、快慢 false | `timeout 5s go run main.go` |
| `nil` | false | map false、快慢 false | `go run main.go` 與臨時 Go 程式 |
| 單節點、無環 | false | false | 臨時 Go 程式 |
| 單節點自環 | true | true | 臨時 Go 程式 |
| 兩節點成環 | true | true | 臨時 Go 程式 |
| 兩節點同值、無環 | false | false | 臨時 Go 程式 |

快慢指標初稿在無環三節點案例曾因讀取 `h2.Next.Next` 時 `h2.Next==nil` 而 panic；修正後上述案例通過。案例通過不等於窮舉證明。

## 面試口述版

用快慢兩指標從 head 出發，每輪分別走一步與兩步，移動後比較是否同一節點。快指標先到尾端表示無環；若有環，進環後快指標相對慢指標每輪多走一步，最終相遇。時間 O(n)、額外空間 O(1)。

## 待複習

- 使用者尚未獨立完整口述 `a+c` 的一般性時間上界，只能指出環內會追上及四節點環的輪數。
- 單節點自環已通過測試，但未獨立口述其停止／相遇流程。
