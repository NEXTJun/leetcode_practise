# 0167. Two Sum II - Input Array Is Sorted - 我的作答

## 作答時間

2026-05-07

## 完整程式碼

```go
package main

import "fmt"

func main() {
	numbers := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum(numbers, target))
}

func twoSum(numbers []int, target int) []int {
	left := 0
	right := len(numbers) - 1

	for left < right {
		num := numbers[left] + numbers[right]
		if target == num {
			return []int{left + 1, right + 1}
		} else if target > num {
			left++
		} else {
			right--
		}
	}

	return nil
}
```

## 你的思路

- 陣列已排序，所以可以用 two pointers。
- `left` 從最小值開始，`right` 從最大值開始。
- 如果總和太小，移動 `left` 讓總和變大。
- 如果總和太大，移動 `right` 讓總和變小。
- 找到時回傳 1-based index。

## 做得好的地方

- 選對方法：排序陣列適合用雙指標。
- 最終版正確處理 1-based index：

```go
return []int{left + 1, right + 1}
```

- 時間複雜度 `O(n)`，空間複雜度 `O(1)`。

## 批改過程中的問題

### 1. 一開始回傳 0-based index

這題要求 1-based index，所以不能回傳：

```go
[]int{left, right}
```

要回傳：

```go
[]int{left + 1, right + 1}
```

### 2. 一開始指標移動方向相反

如果 `sum < target`，代表總和太小，要讓總和變大，因此 `left++`。

如果 `sum > target`，代表總和太大，要讓總和變小，因此 `right--`。

## 可以改進的地方

變數 `num` 可以改成 `sum`，語意更精準：

```go
sum := numbers[left] + numbers[right]
```


## 2026-09-28 面試複習作答

### 原始作答（完整保留）

```go
package main

import "fmt"

func twoSum(numbers []int, target int) []int {
	// TODO: 從陣列兩端開始，依目前的和移動一側指標。
	// 回傳的位置從 1 開始計算。
	var res []int
	var i, j, sum int
	if len(numbers) <= 1 {
		return res
	}
	i = 0
	j = len(numbers) - 1
	sum = numbers[i] + numbers[j]
	res = []int{i + 1, j + 1}

	for sum != target {
		if i >= j {
			res = []int{}
			break
		}
		sum = numbers[i] + numbers[j]
		if sum > target {
			j--
			continue
		} else if sum < target {
			i++
			continue
		}
		res = []int{i + 1, j + 1}
	}
	return res
}

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9)) // 預期：[1 2]
	fmt.Println(twoSum([]int{2, 3, 4}, 6))      // 預期：[1 3]
}
```

### 本輪原本思路與批改

- 使用者先提出以 map 查補數，後能說明排序陣列首尾和太小時應移動左指標。
- 實作中曾將數值與索引混淆，使用 `numbers[len(numbers)]` 導致越界，並先後遇到總和更新、迴圈條件及 0-based 回傳問題；經提示後自行修正。
- 使用者口述：從兩側夾擊，和太大時減小較大數、太小時反向移動；時間 O(n)、額外空間 O(1)；留意 i >= j。
- 排除一側不會漏掉答案的嚴謹理由尚未由使用者獨立說明。
