# 0704. Binary Search - 我的作答

## 作答時間

2026-05-04

## 完整程式碼

```go
package main

import "fmt"

func main() {
	nums := []int{-1, 0, 3, 5, 9, 12}
	target := 9
	fmt.Println(search(nums, target))
}

func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	if len(nums) == 0 {
		return -1
	}
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			return mid
		}
	}

	return -1
}
```

## 你的思路

- 使用左右邊界 `left`、`right` 表示目前搜尋範圍。
- 每次取中間位置 `mid`。
- 如果 `nums[mid]` 等於 `target`，直接回傳 `mid`。
- 如果中間值太大，往左半邊找。
- 如果中間值太小，往右半邊找。

## 做得好的地方

- 選對方法：排序陣列可以用 binary search。
- 使用標準且安全的 mid 寫法：

```go
mid := left + (right-left)/2
```

- 終止條件正確：

```go
for left <= right
```

- 找不到時回傳 `-1`。

## 批改過程中的問題

一開始邊界移動方向與幅度容易搞錯：

```go
if nums[mid] > target {
	left++
} else if nums[mid] < target {
	right--
}
```

這樣有兩個問題：

- 方向反了：中間值太大時，應該縮小右邊界。
- 二分搜尋不是移動一格，而是直接排除半邊。

正確寫法：

```go
if nums[mid] > target {
	right = mid - 1
} else if nums[mid] < target {
	left = mid + 1
}
```

## 可以改進的地方

空陣列檢查可以保留，也可以省略。

保留：

```go
if len(nums) == 0 {
	return -1
}
```

語意清楚。

省略也可以，因為空陣列時 `right := len(nums) - 1` 會是 `-1`，`left <= right` 一開始就是 false，最後自然回傳 `-1`。


## 2026-09-28 面試複習作答

### 原始作答（完整保留）

```go
package main

import "fmt"

func search(nums []int, target int) int {
	// TODO: 在目前的 [left, right] 範圍內取 mid。
	// nums[mid] 小於 target 時更新 left；大於時更新 right。
	// 找到回傳 0-based 索引，找不到回傳 -1。
	res := -1
	if len(nums) == 0 {
		return res
	}
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := (left + right) / 2
		v := nums[mid]
		if v == target {
			res = mid
			break
		} else if v > target {
			right = mid - 1
		} else if v < target {
			left = mid + 1
		}
	}
	return res
}

func main() {
	nums := []int{-1, 0, 3, 5, 9, 12}
	fmt.Println(search(nums, 9)) // 預期：4
	fmt.Println(search(nums, 2)) // 預期：-1
}
```

### 本輪思路與修正

- 使用者指出應從中間搜尋，每輪排除半邊，時間為對數級；對 `nums[mid]<target` 能正確說出 `left=mid+1`。
- 初稿右邊界設為 `len(nums)`、左右移動方向顛倒、使用 `left<right` 漏檢最後一格；經提示修正為閉區間 `[left,right]` 與 `left<=right`。
- 使用者最終口述：排序使被排除半邊的元素全都大於或小於目標；時間 O(log n)、額外空間 O(1)；只剩一格仍需檢查。
