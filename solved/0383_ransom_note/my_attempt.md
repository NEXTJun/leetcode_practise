# 0383. Ransom Note - 我的原始作答

## 作答時間

2026-05-07

## 完整程式碼

```go
package main

import "fmt"

func main() {
	ransomNote := "aa"
	magazine := "aab"
	fmt.Println(canConstruct(ransomNote, magazine))
}

func canConstruct(ransomNote string, magazine string) bool {
	cmap := make(map[rune]int)
	for _, ch := range magazine {
		cmap[ch]++
	}
	for _, ch := range ransomNote {
		cmap[ch]--
	}
	for _, val := range cmap {
		if val < 0 {
			return false
		}
	}
	return true
}
```

## 你的思路

- 統計 `magazine` 裡每個字母可以用幾次。
- 走訪 `ransomNote`，把需要的字母次數扣掉。
- 如果最後有任何字母次數小於 0，代表 `magazine` 不夠組成 `ransomNote`。

## 做得好的地方

- 統計方向正確：應該統計 `magazine`，因為它是可用字母庫。
- 使用 `map[rune]int` 可以處理更通用的字元。
- 邏輯正確，能通過基本測資。

## 可以改進的地方

### 1. 扣字母時可以提早回傳

原本是全部扣完後再掃一次 map：

```go
for _, val := range cmap {
	if val < 0 {
		return false
	}
}
```

可以在扣 `ransomNote` 時直接檢查：

```go
cmap[ch]--
if cmap[ch] < 0 {
	return false
}
```

這樣一旦發現某個字母不夠，就不用繼續走訪。

### 2. 如果只有小寫英文字母，可以用 `[26]int`

`map[rune]int` 較通用。若題目限制只有小寫英文字母，`[26]int` 常數更小，空間也固定為 `O(1)`。
