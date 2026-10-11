# 0001. Two Sum

## 題目

給定整數陣列 `nums` 與整數 `target`，回傳兩個不同 index，使得兩個數字相加等於 `target`。

## 最終作答

```go
func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int)

	for i, val := range nums {
		if j, ok := numMap[target-val]; ok {
			return []int{j, i}
		}
		numMap[val] = i
	}

	return nil
}
```

## 解法重點

- 用 `map[int]int` 記錄「已經看過的數字 -> index」。
- 每次走訪目前數字 `val` 時，先找 `target - val` 是否已經出現過。
- 如果出現過，代表之前的 index `j` 和現在的 index `i` 可以組成答案。
- 查完之後才把目前數字放進 map，避免同一個元素被重複使用。

## 從原始作答改進的地方

- 從「先建立完整 map」改成「一邊走訪，一邊查之前看過的數字」。
- 查到的 complement 一定在目前 index 之前，因此不會用到自己。
- 找到答案就直接 `return`，不需要額外建立 `result` 變數。
- 使用 `j` 或 `idx` 表示 map 中取出的 index，避免把 index 命名成 `val`。

## 複雜度

- 時間複雜度：`O(n)`
- 空間複雜度：`O(n)`

## 面試口述版

我會使用 hash map 儲存已經走訪過的數字與它的 index。對每個數字 `val`，我需要找的是 `target - val`。如果這個 complement 已經在 map 中，代表它是前面出現過的數字，可以和目前 index 組成答案，因此直接回傳。若沒有找到，就把目前數字放入 map，提供後面的元素查詢。這樣每個元素只會走訪一次，所以時間複雜度是 `O(n)`，map 最多存 `n` 個元素，所以空間複雜度是 `O(n)`。
