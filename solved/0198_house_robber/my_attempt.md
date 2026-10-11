# 0198. House Robber - 原始作答

## 2026-10-04 初次完成作答

### 完整原始 `main.go`

```go
package main

import "fmt"

func rob(nums []int) int {
	// TODO: 定義「考慮到目前房子時的最佳金額」，逐間比較拿與不拿。
	// 提醒：拿目前這間時，不能接上相鄰那間的結果。
	nl := len(nums)
	if nl == 0 {
		return 0
	} else if nl == 1 {
		return nums[0]
	}

	F := make([]int, nl)
	for i := 0; i < nl; i++ {
		if i == 0 {
			F[i] = nums[i]
			continue
		} else if i == 1 {
			F[i] = max(nums[i-1], nums[i])
			continue
		}
		F[i] = max(F[i-1], F[i-2]+nums[i])
	}
	return F[nl-1]
}

func max(n1, n2 int) int {
	if n1 > n2 {
		return n1
	}
	return n2
}

func main() {
	// 預期：12、4、5
	fmt.Println(rob([]int{2, 7, 9, 3, 1}))
	fmt.Println(rob([]int{1, 2, 3, 1}))
	fmt.Println(rob([]int{5}))
}
```

### 使用者口述與修正歷程

- 起初用前三間觀察：選 9 時可再選 2，最多 11；之後經例子了解選目前房子須跳過相鄰房子。
- 自行寫出 DP 陣列轉移；前兩次 `check` 因 `i==1` 初始化後仍落入通用轉移，讀取 `F[-1]`，執行時 panic。
- 第三次自行在 `i==1` 分支加入 `continue`，範例和補充案例通過。
- 口述 `F[i]` 為目前最大值、拿時需看 `F[i-2]`、時間 O(n)、空間 O(n)；經提醒後能正確計算 `[2,7,9,3]` 最後一間拿為 10、不拿為 11。
- 尚未獨立完整口述通用正確性證明與邊界案例；已用程式處理空陣列與單間房。
