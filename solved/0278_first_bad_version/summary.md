# 0278. First Bad Version

## 題目

給定版本 `1` 到 `n`，某個版本開始是 bad，之後所有版本也都是 bad。透過 `isBadVersion(version)` API 找出第一個 bad version。

## 最終作答

```go
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
```

## 解法重點

- 版本狀態有單調性：`good good good bad bad bad`。
- 因為有單調性，可以用 binary search。
- 如果 `mid` 是 bad，第一個 bad 可能是 `mid`，所以更新 `right = mid`。
- 如果 `mid` 是 good，第一個 bad 一定在右邊，所以更新 `left = mid + 1`。
- 當 `left == right`，搜尋範圍只剩一個版本，該版本就是第一個 bad。

## 從作答改進的地方

- 修正 `left = mid` 造成區間不縮小的問題。
- 修正 `for left <= right` 搭配 `right = mid` 造成無限迴圈的問題。
- 最終版使用 `for left < right`，和 `right = mid` / `left = mid + 1` 搭配。

## 複雜度

- 時間複雜度：`O(log n)`
- 空間複雜度：`O(1)`

## 面試口述版

這題的版本狀態是單調的：前面都是 good，從某個版本開始都是 bad。所以可以用 binary search。每次檢查中間版本 `mid`，如果 `mid` 是 bad，代表第一個 bad 可能就是 `mid`，也可能在左邊，因此保留 `mid`，更新 `right = mid`。如果 `mid` 是 good，代表答案一定在右邊，所以更新 `left = mid + 1`。當 `left` 和 `right` 相遇時，就找到了第一個 bad version。時間複雜度是 `O(log n)`，空間複雜度是 `O(1)`。
