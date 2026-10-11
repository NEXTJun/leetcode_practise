# 0053. Maximum Subarray

## 題目

給定整數陣列 `nums`，找出總和最大的連續子陣列，回傳該最大總和。

## 最終作答

```go
func maxSubArray(nums []int) int {
	currentSum := nums[0]
	maxSum := nums[0]

	for i := 1; i < len(nums); i++ {
		num := nums[i]

		if currentSum < 0 {
			currentSum = num
		} else {
			currentSum += num
		}

		if currentSum > maxSum {
			maxSum = currentSum
		}
	}

	return maxSum
}
```

## 解法重點

- 這題使用 Kadane's Algorithm。
- `currentSum` 表示「以目前位置結尾」的最大連續子陣列總和。
- `maxSum` 表示目前看過的全域最大總和。
- 如果 `currentSum < 0`，代表前面那段會拖累後面，應該從目前數字重新開始。
- 如果 `currentSum >= 0`，代表前面那段對後續有幫助，可以接上目前數字。
- 每次更新 `currentSum` 後，都要更新 `maxSum`。

## 從原始作答改進的地方

- 原始作答邏輯已正確。
- 最終版將判斷式從 `currentSum + num < num` 簡化成 `currentSum < 0`。
- 從 index `1` 開始走訪，讓初始化和迴圈邏輯更乾淨。
- 保留用 `nums[0]` 初始化，避免全負數時錯誤回傳 `0`。

## 複雜度

- 時間複雜度：`O(n)`
- 空間複雜度：`O(1)`

## 面試口述版

我會用 Kadane's Algorithm。走訪陣列時，用 `currentSum` 記錄目前這段連續子陣列的總和，用 `maxSum` 記錄目前看過的最大總和。每次看到新的數字，如果前面的 `currentSum` 是負數，代表接上它只會讓總和變小，所以從目前數字重新開始；如果 `currentSum` 是非負數，就把目前數字接上去。每次更新完 `currentSum` 後，都用它更新 `maxSum`。因為只走訪一次，所以時間複雜度是 `O(n)`，只用兩個變數，所以空間複雜度是 `O(1)`。


## 2026-09-28 面試複習

- 本輪完整作答原文見 `my_attempt.md`；既有 `final_solution.go` 保留為等價且較簡潔的判斷版本，與上方最終作答對應。
- 正確性：`current` 是所有以當前元素結尾的連續子陣列中最大總和。任何這種子陣列只有「單獨目前元素」或「接在前一格結尾的最佳子陣列後」兩類，因此取兩者較大即可。`best` 保留所有已處理位置的最大 `current`，最後即全域答案。
- 複雜度：`n` 為陣列長度，每個元素處理一次，時間 `O(n)`；只用常數個整數，額外空間 `O(1)`。
- 邊界：輸入保證非空；以 `nums[0]` 初始化可正確處理全負數。若誤以 0 初始化，`[-3,-2,-5]` 會錯回 0，正解是 -2。
- 驗證：`go run main.go` 實際輸出 `6`、`-2`；先前暫存測試中斷後重新執行，單元素正/負、最佳值在前、重新開始與負數夾在最佳區間等六例通過。
- 本輪口述已確認：兩種選擇、O(n)/O(1)、從 index 0 初始化。待複習：用「以目前元素結尾的最大總和」精確定義 current，並獨立說明轉移涵蓋所有連續子陣列。
- 下次複習：2026-10-05。
