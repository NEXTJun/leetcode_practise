# 0217. Contains Duplicate

## 題目

給定整數陣列 `nums`，判斷是否有任何數字出現至少兩次。如果有重複，回傳 `true`；如果每個元素都不同，回傳 `false`。

## 最終作答

```go
func containsDuplicate(nums []int) bool {
	seen := make(map[int]struct{})

	for _, num := range nums {
		if _, ok := seen[num]; ok {
			return true
		}
		seen[num] = struct{}{}
	}

	return false
}
```

## 解法重點

- 用 hash set 記錄已經看過的數字。
- 走訪每個 `num` 時，先檢查它是否存在於 set。
- 如果已經存在，代表有重複，直接回傳 `true`。
- 如果不存在，就把它加入 set。
- 走完整個陣列都沒有發現重複，回傳 `false`。

## 從原始作答改進的地方

- 核心解法完全正確。
- 只把變數命名從 `data` 改成 `seen`，讓語意更清楚。
- 使用 `map[int]struct{}` 表達 set，不需要額外存布林值。

## 複雜度

- 時間複雜度：`O(n)`
- 空間複雜度：`O(n)`

## 面試口述版

我會用 hash set 記錄已經看過的數字。走訪陣列時，如果目前數字已經存在 set 裡，代表它至少出現第二次，可以直接回傳 `true`。如果不存在，就把它加入 set。當整個陣列都走完仍然沒有提早回傳，代表沒有任何重複元素，因此回傳 `false`。時間複雜度是 `O(n)`，因為每個元素只處理一次；空間複雜度是 `O(n)`，因為最壞情況下 set 會存下所有元素。
