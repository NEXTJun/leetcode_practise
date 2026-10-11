# 0035. Search Insert Position

## 題目

給定一個由小到大排序好的整數陣列 `nums` 和整數 `target`。如果 `target` 存在，回傳它的 index；如果不存在，回傳它應該插入的位置，讓陣列仍保持排序。

## 最終作答

```go
func searchInsert(nums []int, target int) int {
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

	return left
}
```

## 解法重點

- 使用 binary search。
- 如果找到 `target`，回傳 `mid`。
- 如果 `nums[mid] < target`，表示 target 應該在右半邊，更新 `left = mid + 1`。
- 如果 `nums[mid] > target`，表示 target 應該在左半邊，更新 `right = mid - 1`。
- 若最後找不到，回傳 `left`。
- 搜尋結束時 `left` 是第一個大於 `target` 的位置，也就是插入位置。

## 從作答改進的地方

- 原始修正版已正確。
- 最終版省略 `num := nums[mid]`，直接使用 `nums[mid]`，減少一個變數。
- 保留 `left + (right-left)/2` 的安全 mid 寫法。

## 複雜度

- 時間複雜度：`O(log n)`
- 空間複雜度：`O(1)`

## 面試口述版

因為陣列已排序，我使用二分搜尋。每次取中間位置 `mid`，如果 `nums[mid]` 等於 `target`，就回傳 `mid`。如果 `nums[mid]` 小於 `target`，代表插入位置或目標值只可能在右半邊，所以更新 `left = mid + 1`；如果比較大，就更新 `right = mid - 1`。當搜尋結束時，`left` 會停在第一個大於 `target` 的位置，所以如果沒找到，`left` 就是正確插入位置。時間複雜度是 `O(log n)`，空間複雜度是 `O(1)`。
