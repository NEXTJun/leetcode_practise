# 0383. Ransom Note

## 題目

給定 `ransomNote` 和 `magazine` 兩個字串，判斷 `ransomNote` 能不能由 `magazine` 裡的字母組成。`magazine` 裡每個字母只能使用一次。

## 最終作答：map 版本

```go
func canConstruct(ransomNote string, magazine string) bool {
	counts := make(map[rune]int)

	for _, ch := range magazine {
		counts[ch]++
	}

	for _, ch := range ransomNote {
		counts[ch]--
		if counts[ch] < 0 {
			return false
		}
	}

	return true
}
```

## 解法重點

- `magazine` 是可用字母庫，所以先統計 `magazine` 的字母次數。
- 走訪 `ransomNote` 時，逐一扣掉需要的字母。
- 如果某個字母被扣到負數，代表 `magazine` 不夠用，直接回傳 `false`。
- 如果全部扣完都沒有負數，代表可以組成，回傳 `true`。

## 另一個最優版本：小寫英文字母計數

如果題目限制只有 `a-z`，可以用 `[26]int`：

```go
func canConstructCounting(ransomNote string, magazine string) bool {
	var counts [26]int

	for i := 0; i < len(magazine); i++ {
		counts[magazine[i]-'a']++
	}

	for i := 0; i < len(ransomNote); i++ {
		idx := ransomNote[i] - 'a'
		counts[idx]--
		if counts[idx] < 0 {
			return false
		}
	}

	return true
}
```

## 從原始作答改進的地方

- 原始作答邏輯正確。
- 最終版將變數 `cmap` 改成 `counts`，語意更清楚。
- 在扣 `ransomNote` 時提早檢查負數，避免最後再掃一次 map。
- 補充 `[26]int` 版本，方便面試時說明小寫限制下的空間最佳化。

## 複雜度

map 版本：

- 時間複雜度：`O(n + m)`
- 空間複雜度：`O(k)`

`[26]int` 版本：

- 時間複雜度：`O(n + m)`
- 空間複雜度：`O(1)`

其中 `n = len(magazine)`，`m = len(ransomNote)`，`k` 是不同字元種類數。

## 面試口述版

我會先統計 `magazine` 中每個字母的出現次數，因為它代表我手上可用的字母庫。接著走訪 `ransomNote`，每需要一個字母就把對應次數減一。如果某個字母被扣到負數，代表 `magazine` 裡這個字母不夠用，直接回傳 `false`。如果整個 `ransomNote` 都能成功扣完，就回傳 `true`。時間複雜度是 `O(n + m)`。如果用 map，空間是 `O(k)`；若題目限制只有小寫英文字母，可以用長度 26 的陣列把空間壓到 `O(1)`。
