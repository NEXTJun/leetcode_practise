# 0001. Two Sum - 我的原始作答

## 作答時間

2026-05-03

## 完整程式碼

```go
package main

import "fmt"

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9

	fmt.Println(twoSum(nums, target))
}

func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int)
	for i, val := range nums {
		numMap[val] = i
	}
	var result []int
	for j, val := range nums {
		val, ok := numMap[target-val]
		if ok {
			result = []int{j, val}
			break
		}
	}
	return result
}
```

## 你的思路

- 先用 `map[int]int` 建立「數字 -> index」的查找表。
- 再走訪一次 `nums`，用 `target - val` 找另一個數字是否存在。
- 找到後回傳目前 index 與 map 中存的 index。

## 做得好的地方

- 已經想到用 hash map，把查找從線性搜尋降成近似 `O(1)`。
- 整體方向比暴力解好，時間複雜度目標是 `O(n)`。

## 需要修正的地方

### 1. 可能使用同一個元素兩次

因為你先把所有元素都放進 map，所以查詢時可能查到自己。

反例：

```go
nums := []int{3, 2, 4}
target := 6
```

走到 index `0`、值 `3` 時，會查 `target - val = 3`，map 裡找到的也是 index `0`，可能得到：

```text
[0 0]
```

這違反題目規則：同一個元素不能使用兩次。

### 2. 變數命名覆蓋了原本的 `val`

這段：

```go
val, ok := numMap[target-val]
```

左邊新的 `val` 其實變成 index，不再是原本 nums 裡的數值。雖然 Go 可以編譯，但閱讀上容易混淆。

建議改成：

```go
idx, ok := numMap[target-val]
```

### 3. 找到答案可以直接回傳

現在使用 `result` 加 `break`：

```go
result = []int{j, val}
break
```

因為題目保證有唯一答案，找到時可以直接：

```go
return []int{j, i}
```

程式會更短，也更符合面試解法表達。
