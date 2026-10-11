# 0125. Valid Palindrome

## 題目

給定字串 `s`，判斷它是否是 palindrome。只考慮英文字母與數字，忽略大小寫與其他符號。

## 最終作答

```go
func isPalindrome(s string) bool {
	left := 0
	right := len(s) - 1

	for left < right {
		for left < right && !isAlphaNum(s[left]) {
			left++
		}
		for left < right && !isAlphaNum(s[right]) {
			right--
		}

		if toLower(s[left]) != toLower(s[right]) {
			return false
		}

		left++
		right--
	}

	return true
}

func isAlphaNum(ch byte) bool {
	return (ch >= 'a' && ch <= 'z') ||
		(ch >= 'A' && ch <= 'Z') ||
		(ch >= '0' && ch <= '9')
}

func toLower(ch byte) byte {
	if ch >= 'A' && ch <= 'Z' {
		return ch - 'A' + 'a'
	}
	return ch
}
```

## 解法重點

- 使用 two pointers：`left` 從左往右，`right` 從右往左。
- 不建立新字串，直接在原字串上跳過非英數字元。
- 比較前將字母轉成小寫。
- 若任一組有效字元不同，直接回傳 `false`。
- 指標交錯或相遇時代表檢查完成，回傳 `true`。

## 從原始作答改進的地方

- 原始作答已正確通過邊界測資。
- 最終版移除 `lb` / `rb` 暫存變數，讓流程更短。
- 讓 `left` 和 `right` 停在有效字元後直接比較，面試口述更清楚。
- `isAlphaNum` 可直接回傳布林表達式。

## 複雜度

- 時間複雜度：`O(n)`
- 空間複雜度：`O(1)`

## 面試口述版

我會用雙指標從字串左右兩端往中間檢查。因為題目只要求英數字元，所以如果左指標或右指標遇到非英數字元，就直接跳過。當兩邊都停在有效字元時，把字母轉成小寫後比較；如果不同，代表不是 palindrome，回傳 `false`。如果一路檢查到兩個指標相遇或交錯都沒有失敗，就回傳 `true`。整個字串最多被走訪一次，所以時間複雜度是 `O(n)`，只用常數變數，所以空間複雜度是 `O(1)`。
