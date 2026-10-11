# 0242. Valid Anagram - 我的原始作答

## 作答時間

2026-05-04

## 完整程式碼

```go
package main

import "fmt"

func main() {
	s := "anagram"
	t := "nagaram"
	fmt.Println(isAnagram(s, t))
}

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	var words [26]int

	for _, ch := range s {
		words[ch-'a']++
	}

	for _, ch := range t {
		words[ch-'a']--
	}

	for _, num := range words {
		if num != 0 {
			return false
		}
	}

	return true
}
```

## 你的思路

- 先檢查兩個字串長度是否相同。
- 使用 `[26]int` 記錄小寫英文字母出現次數。
- 走訪 `s` 時，對應字母計數加一。
- 走訪 `t` 時，對應字母計數減一。
- 最後檢查所有計數是否歸零。

## 做得好的地方

- 解法正確。
- 有先處理長度不同的情況。
- 使用 `[26]int` 很適合本題小寫英文字母限制。
- 時間複雜度是 `O(n)`，空間複雜度是 `O(1)`。

## 可以改進的地方

### 1. 變數命名可以更精準

原本：

```go
var words [26]int
```

建議：

```go
var counts [26]int
```

因為陣列中存的是每個字母的次數，不是單字。

### 2. 可在扣次數時提早判斷

若先統計 `s`，再走訪 `t` 扣次數，當某個字母被扣到負數時，代表 `t` 中該字母比 `s` 多，可以直接回傳 `false`。

### 3. 排序法也要會

本題另一種常見解法是把兩個字串排序後比較。排序法比較直觀，也較容易延伸到不是固定 26 個小寫字母的情境。
