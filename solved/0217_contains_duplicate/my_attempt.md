# 0217. Contains Duplicate - 我的原始作答

## 作答時間

2026-05-04

## 完整程式碼

```go
package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 1}
	fmt.Println(containsDuplicate(nums))
}

func containsDuplicate(nums []int) bool {
	data := make(map[int]struct{})
	for _, num := range nums {
		if _, ok := data[num]; ok {
			return true
		}
		data[num] = struct{}{}
	}
	return false
}
```

## 你的思路

- 使用 hash set 記錄已經看過的數字。
- 每次走訪一個 `num`，先檢查它是否已經存在 set 裡。
- 如果存在，代表有重複，直接回傳 `true`。
- 如果不存在，就把它加入 set。
- 全部走完都沒有重複，就回傳 `false`。

## 做得好的地方

- 選對資料結構：hash set。
- 使用 Go 裡標準的 set 寫法：

```go
map[int]struct{}
```

- 遇到重複就立刻 `return true`，不做多餘走訪。
- 時間複雜度與空間複雜度都分析清楚。

## 可以改進的地方

### 1. 命名可以更精準

原本：

```go
data := make(map[int]struct{})
```

可以改成：

```go
seen := make(map[int]struct{})
```

`seen` 更能表達「已經看過的數字」。

### 2. 解法已經正確

這題沒有核心邏輯錯誤。只需要在面試時說清楚為什麼使用 set，以及為什麼可以提早回傳。
