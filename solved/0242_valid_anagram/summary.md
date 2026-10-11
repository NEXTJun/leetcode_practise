# 0242. Valid Anagram

## 題目

給定兩個字串 `s` 和 `t`，判斷 `t` 是否是 `s` 的 anagram。Anagram 代表兩個字串使用的字元與每個字元出現次數都相同，只是順序可以不同。

## 方法一：排序法

排序法的想法最直觀：如果兩個字串是 anagram，排序後的字元順序會完全相同。

```go
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	bs := []byte(s)
	bt := []byte(t)

	sort.Slice(bs, func(i, j int) bool {
		return bs[i] < bs[j]
	})
	sort.Slice(bt, func(i, j int) bool {
		return bt[i] < bt[j]
	})

	return string(bs) == string(bt)
}
```

複雜度：

- 時間複雜度：`O(n log n)`
- 空間複雜度：`O(n)`

適合情境：

- 想法直觀，容易先在面試中說明。
- 字元種類不固定時，比 `[26]int` 更容易泛化。

## 方法二：計數法

如果題目限制只有小寫英文字母 `a-z`，可以用長度固定的 `[26]int` 統計字母次數。

```go
func isAnagramCounting(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	var counts [26]int

	for _, ch := range s {
		counts[ch-'a']++
	}

	for _, ch := range t {
		counts[ch-'a']--
	}

	for _, count := range counts {
		if count != 0 {
			return false
		}
	}

	return true
}
```

複雜度：

- 時間複雜度：`O(n)`
- 空間複雜度：`O(1)`

適合情境：

- 題目明確限制只有小寫英文字母。
- 面試官追問最優時間複雜度時，這是更好的答案。

## 從原始作答改進的地方

- 原始作答使用 `[26]int` 計數，核心解法正確。
- 將變數名稱從 `words` 建議改成 `counts`，語意更精準。
- 補充排序法作為另一種常見解法，方便面試時先講直觀方法，再優化到計數法。

## 取捨

- 排序法：直觀、好解釋、容易泛化；但時間是 `O(n log n)`。
- 計數法：在小寫英文字母限制下複雜度最好；時間 `O(n)`、空間 `O(1)`。
- 面試建議：先能講排序法，再主動補充「因為本題只有 `a-z`，可以用 `[26]int` 優化」。

## 面試口述版

我會先檢查兩個字串長度是否相同，若不同就不可能是 anagram。直觀做法是把兩個字串排序，因為 anagram 的字元組成和次數相同，所以排序後應該完全一樣，這個方法是 `O(n log n)`。如果題目限制只有小寫英文字母，我會進一步用長度 26 的陣列統計字母次數：走訪 `s` 時加一，走訪 `t` 時減一，最後所有計數都為 0 就代表是 anagram。這個計數法時間是 `O(n)`，空間是 `O(1)`。
