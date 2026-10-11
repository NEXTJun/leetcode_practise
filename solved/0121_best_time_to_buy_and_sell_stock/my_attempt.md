# 0121. Best Time to Buy and Sell Stock - 我的原始作答

## 作答時間

2026-05-03

## 完整程式碼

```go
package main

import "fmt"

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(prices))
}

func maxProfit(prices []int) int {
	var minPrice, maxProfit int
	if len(prices) == 0 {
		return 0
	}
	minPrice = prices[0]
	for _, price := range prices {
		if price-minPrice > maxProfit {
			maxProfit = price - minPrice
		}
		if price < minPrice {
			minPrice = price
		}
	}
	return maxProfit
}
```

## 你的思路

- 用 `minPrice` 記錄目前走過的最低買入價格。
- 用 `maxProfit` 記錄目前可以得到的最大利潤。
- 每天都先假設今天賣出，計算 `price - minPrice`。
- 如果今天價格更低，就更新 `minPrice` 給後面的日期使用。

## 做得好的地方

- 解法正確，而且是一次走訪。
- 有處理空陣列：

```go
if len(prices) == 0 {
	return 0
}
```

- 只用兩個變數，空間複雜度是 `O(1)`。
- 更新順序合理：先用目前最低價計算今天賣出的利潤，再更新今天是否成為新的最低價。

## 可以改進的地方

### 1. 初始化可以更直接

原本寫法：

```go
var minPrice, maxProfit int
minPrice = prices[0]
```

可以改成：

```go
minPrice := prices[0]
maxProfit := 0
```

語意更直接，也少一行。

### 2. 可讀性可以用 `profit` 暫存

原本寫法：

```go
if price-minPrice > maxProfit {
	maxProfit = price - minPrice
}
```

可以改成：

```go
profit := price - minPrice
if profit > maxProfit {
	maxProfit = profit
}
```

這樣面試時更容易對照口述邏輯。

## 2026-09-29 複習作答

### 原始作答（完整 `main.go`）

```go
package main

import "fmt"

func maxProfit(prices []int) int {
	// TODO: 從左到右走訪，記錄此前最低買價與目前最佳利潤。
	// 想想：今天的價格要先拿來算賣出利潤，還是先更新最低買價？
	var low, maxP int
	if len(prices) == 0 {
		return 0
	}
	for i := 0; i < len(prices); i++ {
		if i == 0 {
			low = prices[i]
			continue
		}
		if prices[i] > low {
			if prices[i]-low > maxP {
				maxP = prices[i] - low
			}
		} else {
			low = prices[i]
		}
	}
	return maxP
}

func main() {
	// 預期：5、0、0
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
	fmt.Println(maxProfit([]int{7, 6, 4, 3, 1}))
	fmt.Println(maxProfit([]int{5}))
}
```

### 本輪思路與檢查

- 先提出逐日固定買入日，再掃描未來最高價，這是正確的 O(n²) 暴力法。
- 經提示轉成逐日考慮賣出價；使用者能算出價格 6 時此前最低 1、當日利潤 5、歷史最佳 5。
- 使用者自行寫出維護過去最低價與最佳利潤的一次走訪實作，並口述先買後賣、逐日與歷史最大利潤比較，以及 O(n) 時間、O(1) 額外空間。
- 本輪沒有發現需修正的邏輯錯誤；scaffold 留下的 TODO 註解不影響執行。
- `go run main.go` 實際輸出 5、0、0；另執行 `[2,4,1]`、`[3,3,5,0,0,3,1,4]`、`[]`，實際輸出 2、4、0，均符合預期。
