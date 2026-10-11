# 0035. Search Insert Position - 我的作答

## 作答時間

2026-05-07

## 完整程式碼

```go
package main

import "fmt"

func main() {
	nums := []int{1, 3, 5, 6}
	target := 2
	fmt.Println(searchInsert(nums, target))
}

func searchInsert(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := left + (right-left)/2
		num := nums[mid]
		if num == target {
			return mid
		} else if num > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return left
}
```

## 你的思路

- 這題和 Binary Search 類似。
- 如果找到 `target`，直接回傳 `mid`。
- 如果找不到，最後回傳 `left`。
- 因為搜尋結束時，`left` 會停在第一個大於 `target` 的位置，也就是插入位置。

## 做得好的地方

- 最終版使用正確的二分搜尋邊界：

```go
for left <= right
```

- 找不到時正確回傳 `left`。
- 使用安全 mid 寫法：

```go
mid := left + (right-left)/2
```

## 批改過程中的問題

一開始把 index 和 value 混在一起：

```go
num := nums[mid]
right = num
left = num + 1
```

這會造成邊界更新錯誤，甚至可能無限迴圈。

二分搜尋裡要記住：

```text
left / right / mid 是 index
nums[mid] 是 value
```

所以邊界更新必須用：

```go
right = mid - 1
left = mid + 1
```

## 可以改進的地方

目前解法已正確。面試時要特別講清楚為什麼找不到時回傳 `left`：

- `right` 停在最後一個小於 `target` 的位置。
- `left` 停在第一個大於 `target` 的位置。
- 插入位置應該在這兩者中間，也就是 `left`。
