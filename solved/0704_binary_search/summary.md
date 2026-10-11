# 0704. Binary Search

## 題目

給定一個由小到大排序好的整數陣列 `nums` 和整數 `target`。如果 `target` 存在，回傳它的 index；如果不存在，回傳 `-1`。

## 最終作答

```go
func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := left + (right-left)/2

		if nums[mid] == target {
			return mid
		}
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}
```

## 解法重點

- 陣列已排序，所以可以每次排除一半搜尋範圍。
- `left` 和 `right` 表示目前可能包含答案的區間。
- `mid` 是目前區間中點。
- 如果 `nums[mid] == target`，直接回傳 `mid`。
- 如果 `nums[mid] < target`，代表 target 只可能在右半邊，更新 `left = mid + 1`。
- 如果 `nums[mid] > target`，代表 target 只可能在左半邊，更新 `right = mid - 1`。
- 當 `left > right`，代表搜尋區間為空，回傳 `-1`。

## 從作答改進的地方

- 最終版省略空陣列特判，因為迴圈條件已能自然處理空陣列。
- 將判斷順序整理成先判斷相等，再處理左右邊界，口述更直覺。
- 核心修正是邊界更新必須使用 `mid + 1` 與 `mid - 1`，不是 `left++` 或 `right--`。

## 複雜度

- 時間複雜度：`O(log n)`
- 空間複雜度：`O(1)`

## 面試口述版

因為陣列已經排序，我可以用二分搜尋。用 `left` 和 `right` 表示目前搜尋範圍，每次取中間 index `mid`。如果 `nums[mid]` 等於 target，就直接回傳 `mid`；如果 `nums[mid]` 比 target 小，代表 target 只可能在右半邊，所以更新 `left = mid + 1`；如果 `nums[mid]` 比 target 大，代表 target 只可能在左半邊，所以更新 `right = mid - 1`。當 `left > right` 時代表找不到，回傳 `-1`。每次都排除一半範圍，所以時間複雜度是 `O(log n)`，只使用常數變數，所以空間複雜度是 `O(1)`。


## 2026-09-28 面試複習

- 本輪完整作答原文見 `my_attempt.md`。既有 `final_solution.go` 為等價的簡潔版本，與上方最終作答對應；本輪程式保留空陣列特判與 `res` 變數。
- 正確性：閉區間 `[left,right]` 包含尚可能是答案的索引。若中點值小於目標，排序保證中點及左側都不可能是答案，故 `left=mid+1`；反之 `right=mid-1`。相等時回傳索引，區間空時回傳 -1。
- 複雜度：每輪搜尋範圍至少約減半，時間 `O(log n)`；只用固定數量的整數，額外空間 `O(1)`。
- 邊界：閉區間只剩一格時仍要檢查，因此條件為 `left<=right`。找不到時最後變成 `left>right`。
- 驗證：`go run main.go` 實際輸出 `4`、`-1`；暫存 Go 測試的單元素命中/未命中、首尾元素、區間內/外未命中六例通過。
- 本輪口述已確認：排序排除理由、O(log n)/O(1)、最後一格邊界。
- 下次複習：2026-10-05。
