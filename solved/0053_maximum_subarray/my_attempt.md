# 0053. Maximum Subarray - 我的原始作答

## 作答時間

2026-05-07

## 完整程式碼

```go
package main

import "fmt"

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(maxSubArray(nums))
}

func maxSubArray(nums []int) int {
	var currentSum, maxSum int
	for i, num := range nums {
		if i == 0 {
			currentSum = num
			maxSum = num
			continue
		}
		if currentSum+num < num {
			currentSum = num
		} else {
			currentSum = currentSum + num
		}
		if currentSum > maxSum {
			maxSum = currentSum
		}
	}
	return maxSum
}
```

## 你的思路

- `currentSum` 表示目前這段連續子陣列的總和。
- `maxSum` 表示目前看過的最大連續子陣列總和。
- 每次遇到新的 `num`，判斷「接上前一段」還是「從目前數字重新開始」。
- 如果 `currentSum + num < num`，代表接上前一段比較差，改從 `num` 重新開始。
- 每輪都更新 `maxSum`。

## 做得好的地方

- 有正確處理全負數情境，沒有把 `maxSum` 初始化成 `0`。
- 使用第一個元素初始化，方向正確。
- 修掉了重設 `currentSum` 後沒有更新 `maxSum` 的問題。
- 時間複雜度 `O(n)`，空間複雜度 `O(1)`。

## 批改過程中的問題

一開始重設 `currentSum` 後直接 `continue`：

```go
if currentSum+num < num {
	currentSum = num
	continue
}
```

這會漏掉更新 `maxSum`。

反例：

```go
nums := []int{-2, 1}
```

正確答案是 `1`，但如果重設後直接 `continue`，`maxSum` 會停在 `-2`。

## 可以改進的地方

這段：

```go
if currentSum+num < num {
	currentSum = num
} else {
	currentSum = currentSum + num
}
```

可以改成更好口述的版本：

```go
if currentSum < 0 {
	currentSum = num
} else {
	currentSum += num
}
```

因為 `currentSum + num < num` 等價於 `currentSum < 0`。


## 2026-09-28 面試複習作答

### 原始作答（完整保留）

```go
package main

import "fmt"

func maxSubArray(nums []int) int {
	// TODO: current 表示必須以目前元素結尾的最大總和。
	// TODO: best 表示目前看過的最大總和。
	// nums 非空；注意全為負數的案例。
	var current, best int
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			current = nums[i]
			best = current
			continue
		}
		if current+nums[i] >= nums[i] {
			current = current + nums[i]
		} else {
			current = nums[i]
		}
		if current > best {
			best = current
		}
	}
	return best
}

func main() {
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})) // 預期：6
	fmt.Println(maxSubArray([]int{-3, -2, -5}))                    // 預期：-2
}
```

### 本輪思路與修正

- 使用者起初嘗試逐一累加、排序及捨去負數；經反例理解連續子陣列不能重新排序或跳過中間元素。
- 對 `current` 與 `best` 一度混淆，經 [5,-6,4]、[2,-1,3] 等例子能算出續接與重開後的值。
- 初稿曾將 current+nums[i] 與 best 比較、在重開分支漏更新 best 或無條件覆蓋 best；最後改為先選 current，再以 current 更新 best。
- 使用者口述能說明續接／重開判斷、O(n) 時間、O(1) 額外空間及非空限制下從 nums[0] 初始化；current 的精確定義仍可加強。
