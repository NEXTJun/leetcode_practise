# 0278. First Bad Version - 我的作答

## 作答時間

2026-05-07

## 完整程式碼

```go
package main

import "fmt"

var badVersion = 4

func main() {
	n := 5
	fmt.Println(firstBadVersion(n))
}

func firstBadVersion(n int) int {
	left := 1
	right := n

	for left < right {
		mid := left + (right-left)/2
		if isBadVersion(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left
}

func isBadVersion(version int) bool {
	return version >= badVersion
}
```

## 你的思路

- 版本狀態有順序性：`good good good bad bad bad`。
- 使用 binary search 找第一個 bad。
- 如果 `mid` 是 bad，第一個 bad 可能就是 `mid`，也可能在左邊。
- 如果 `mid` 是 good，第一個 bad 一定在右邊。

## 做得好的地方

- 最終版使用正確的二分搜尋方向。
- `mid` 是 bad 時保留 `mid`：

```go
right = mid
```

- `mid` 是 good 時排除 `mid`：

```go
left = mid + 1
```

- 使用 `left + (right-left)/2` 避免溢位風險。

## 批改過程中的問題

### 1. `left = mid` 會卡住

如果 `mid` 是 good，答案一定在右邊，所以要寫：

```go
left = mid + 1
```

不能寫：

```go
left = mid
```

否則在 `left` 和 `right` 相鄰時，`mid` 可能等於 `left`，範圍不會縮小。

### 2. `for left <= right` 搭配 `right = mid` 會卡住

當 `left == right` 且該位置是 bad：

```go
right = mid
```

不會讓範圍變小，會造成無限迴圈。

這種寫法要搭配：

```go
for left < right
```

## 可以改進的地方

目前解法已正確。這題要特別記住：

- 一般找值的 binary search 常用 `left <= right`。
- 找第一個符合條件的位置時，常用 `left < right`。
- 當 `mid` 可能是答案時，不要丟掉 `mid`。
