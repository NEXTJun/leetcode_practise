# 0070. Climbing Stairs

## 題目

你要爬到第 `n` 階樓梯，每次可以爬 `1` 階或 `2` 階。回傳到達第 `n` 階共有幾種不同方法。

## 方法一：Memoization

```go
func climbStairsMemo(n int) int {
	memo := make(map[int]int)

	var dfs func(int) int
	dfs = func(step int) int {
		if step == 1 {
			return 1
		}
		if step == 2 {
			return 2
		}
		if val, ok := memo[step]; ok {
			return val
		}

		ways := dfs(step-1) + dfs(step-2)
		memo[step] = ways
		return ways
	}

	return dfs(n)
}
```

複雜度：

- 時間複雜度：`O(n)`
- 空間複雜度：`O(n)`

## 方法二：Rolling Variables

```go
func climbStairs(n int) int {
	if n <= 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}

	prev2 := 1
	prev1 := 2

	for i := 3; i <= n; i++ {
		current := prev1 + prev2
		prev2 = prev1
		prev1 = current
	}

	return prev1
}
```

複雜度：

- 時間複雜度：`O(n)`
- 空間複雜度：`O(1)`

## 解法重點

- 到第 `n` 階的最後一步，只可能從：
  - `n-1` 階走 1 階
  - `n-2` 階走 2 階
- 所以遞推式是：

```text
ways[n] = ways[n-1] + ways[n-2]
```

- `n = 1` 時答案是 `1`。
- `n = 2` 時答案是 `2`。
- 因為每次只需要前兩個狀態，最終可以用兩個變數壓縮空間。

## 從作答改進的地方

- 原始最終作答是 memoization，邏輯正確。
- 紀錄中補上 rolling variables 版本，這是空間更優的解法。
- 面試時可以先講 memoization，再補充因為只依賴前兩項，所以能優化成 `O(1)` 空間。

## 面試口述版

這題的最後一步只可能從第 `n-1` 階走 1 階，或從第 `n-2` 階走 2 階，所以 `ways[n] = ways[n-1] + ways[n-2]`。一開始可以用 recursion 加 memoization，把每個 `n` 的答案存起來，避免重複計算，時間是 `O(n)`，空間是 `O(n)`。進一步觀察到每次只需要前兩階答案，所以可以用兩個變數 `prev2` 和 `prev1` 滾動更新，把空間降到 `O(1)`。


## 2026-09-29 面試複習

- 本輪完整作答原文見 `my_attempt.md`。既有 `final_solution.go` 保留 rolling 與 memo 兩種複習版本；本輪使用者獨立完成的是 map memo 版，不應把 rolling 版本算成本輪已實作。
- 正確性：任何到第 `n` 階的走法，最後一步不是從 `n-1` 走 1 階，就是從 `n-2` 走 2 階。兩組涵蓋全部走法且不重疊，故 `W(n)=W(n-1)+W(n-2)`，基本情況 `W(1)=1`、`W(2)=2`。使用者能在具體例子中套用並推廣到增加 3 階選項；完整證明主要由教學整理。
- 複雜度：memo 中不同階數最多 `n` 個，每個只計算一次，平均時間 `O(n)`；map 與最深遞迴堆疊各最多 `O(n)`，合計額外空間 `O(n)`。rolling 版因只依賴前兩項可降為 `O(1)` 空間，但本輪未獨立實作。
- 邊界：題目要求正整數 `n`，`n=1` 與 `n=2` 直接回傳 1 與 2。現有程式對 `n=0` 回傳 0；這不影響原題限制內的結果。若討論可走 1、2、3 階的通用遞推，需另外定義 `W(0)=1` 才能把空路線算入。
- 驗證：最新 `go run main.go` 實際輸出 `1`、`2`、`8`；較大階數暫存測試曾被中斷，未列為通過。此前 memo 初稿的 stack overflow 已由實際執行觀察到，最終版本沒有再出現。
- 待複習：不看提示完整口述兩組走法為何不重疊、遞迴堆疊 O(n) 的理由，以及自己實作 rolling O(1) 版本。
- 下次複習：2026-10-02。
