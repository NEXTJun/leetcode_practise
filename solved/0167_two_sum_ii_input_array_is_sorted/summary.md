# 0167. Two Sum II - Input Array Is Sorted

## 題目

給定已排序的整數陣列 `numbers` 和整數 `target`，找出兩個數字相加等於 `target`，回傳它們的 1-based index。

## 最終作答

```go
func twoSum(numbers []int, target int) []int {
	left := 0
	right := len(numbers) - 1

	for left < right {
		sum := numbers[left] + numbers[right]

		if sum == target {
			return []int{left + 1, right + 1}
		}
		if sum < target {
			left++
		} else {
			right--
		}
	}

	return nil
}
```

## 解法重點

- 陣列已排序，可以用 two pointers。
- `left` 指向目前較小的數，`right` 指向目前較大的數。
- 如果 `sum == target`，回傳 1-based index。
- 如果 `sum < target`，代表總和太小，移動 `left++`。
- 如果 `sum > target`，代表總和太大，移動 `right--`。

## 從作答改進的地方

- 原始修正版已正確。
- 最終版把 `num` 改成 `sum`，讓變數語意更清楚。
- 保留 `return nil` 作為防守；LeetCode 題目通常保證有唯一答案。

## 複雜度

- 時間複雜度：`O(n)`
- 空間複雜度：`O(1)`

## 面試口述版

因為陣列已經排序，我可以用左右雙指標。`left` 從最小值開始，`right` 從最大值開始。每次計算兩者總和，如果等於 `target`，就回傳它們的 1-based index。如果總和太小，代表需要更大的數，所以移動 `left++`；如果總和太大，代表需要更小的數，所以移動 `right--`。每次至少移動一個指標，因此時間複雜度是 `O(n)`，只用常數變數，所以空間複雜度是 `O(1)`。


## 2026-09-28 面試複習

- 本輪作答原文見 `my_attempt.md`；原有 `final_solution.go` 保留為簡潔複習版，與上方最終作答對應。
- 本輪修正：左右指標改用索引、右指標設為 `len(numbers)-1`、在移動後重新計算總和，兩處回傳均轉為 1-based。
- 驗證：`go run main.go` 實際輸出 `[1 2]`、`[1 3]`；暫存 Go 測試對 `[-1,1], target=0`、`[2,7], target=9`、`[2,3,4], target=6`、`[2,2], target=4`、`[-4,-1,0,3,10], target=9` 均通過。測試不等於窮舉證明。
- 正確性補充（教學整理，尚非使用者獨立口述）：若 `numbers[left]+numbers[right] < target`，因排序，`numbers[left]` 與任何仍在區間內的元素相加都不會更大，故可排除 `left`。若總和過大，`numbers[right]` 與任何仍在區間內的元素相加都不會更小，故可排除 `right`。有解時不會排除唯一答案。
- 本輪已能口述：雙指標方向、時間 `O(n)`、額外空間 `O(1)`、指標相遇邊界。
- 待複習：用排序性質明確說出兩側排除論證；目前由使用者的口述只能確認直觀移動方向。
- 下次複習：2026-10-05。
