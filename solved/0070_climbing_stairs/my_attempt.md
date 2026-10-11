# 0070. Climbing Stairs - 我的作答

## 作答時間

2026-05-10

## 完整程式碼

```go
package main

import "fmt"

func main() {
	n := 3
	fmt.Println(climbStairs(n))
}

func climbStairs(n int) int {
	memo := make(map[int]int)
	var dfs func(int) int
	dfs = func(step int) int {
		if step == 1 {
			return 1
		}
		if step == 2 {
			return 2
		}

		if val, ok := memo[step]; ok {
			return val
		}
		way := dfs(step-1) + dfs(step-2)
		memo[step] = way
		return way
	}
	return dfs(n)
}
```

## 你的思路

- 到第 `n` 階的最後一步，只可能從 `n-1` 或 `n-2` 來。
- 遞推式是：

```text
ways[n] = ways[n-1] + ways[n-2]
```

- 先想到純 recursion，再用 map 做 memoization 避免重複計算。

## 做得好的地方

- 能看出這題是 Fibonacci 型問題。
- 能想到用 map 記錄已算過的答案。
- 最終 memoization 寫法正確：

```go
way := dfs(step-1) + dfs(step-2)
memo[step] = way
```

## 批改過程中的問題

### 1. 在 helper 裡呼叫外層函式

一開始寫成：

```go
way := climbStairs(n-1) + climbStairs(n-2)
```

這會每次重新建立新的 `memo`，所以 memoization 不會生效。

### 2. 在 helper 裡使用外層固定的 `n`

後來寫成：

```go
way := dfs(n-1) + dfs(n-2)
```

但 `n` 是外層固定值，不會隨遞迴改變，會導致無限遞迴。

在 `dfs(step int)` 裡，base case 和遞迴都要使用 `step`：

```go
if step == 1 { return 1 }
if step == 2 { return 2 }
way := dfs(step-1) + dfs(step-2)
```

## 可以改進的地方

memoization 版本時間是 `O(n)`，空間是 `O(n)`。這題還可以用 rolling variables，只保留前兩階答案，把空間降成 `O(1)`。


## 2026-09-29 面試複習作答

### 原始作答（完整保留）

```go
package main

import "fmt"

func climbStairs(n int) int {
	mapClimb := make(map[int]int)
	// TODO: 先處理 n=1、n=2。
	// TODO: 從第 3 階起，用前兩階的走法數求目前階。
	return climbCal(n, mapClimb)
}

func climbCal(n int, mapClimb map[int]int) int {
	val, ok := mapClimb[n]
	if ok {
		return val
	}
	var res int
	switch n {
	case 0:
		res = 0
	case 1:
		res = 1
	case 2:
		res = 2
	default:
		res = climbCal(n-1, mapClimb) + climbCal(n-2, mapClimb)
	}
	mapClimb[n] = res
	return res
}

func main() {
	fmt.Println(climbStairs(1)) // 預期：1
	fmt.Println(climbStairs(2)) // 預期：2
	fmt.Println(climbStairs(5)) // 預期：8
}
```

### 本輪思路、修正與口述

- 使用者先將到第 4 階的最後一步拆成從第 3 階走 1 階或從第 2 階走 2 階；經具體路線分組後，能給出 W(5)=W(4)+W(3)=8，並主動推廣到可走 1、2、3 階時的三項遞推。
- 初版無記憶遞迴對 n=1、2、5 範例通過，但會重算同一階；改寫 map memo 時曾讓基底 case 落入後續遞迴，實際造成 stack overflow；又曾在 helper 中呼叫外層函式而每次重建 map。最後改用同一 helper 與 map，並把遞迴放入 switch 的 default。
- 使用者口述：map 存已算走法數，所以每階只需真正計算一次，時間 O(n)；map 佔 O(n) 空間。遞迴堆疊 O(n) 由教學補充，未確認使用者能獨立說明；rolling O(1) 版本未在本輪獨立實作。
