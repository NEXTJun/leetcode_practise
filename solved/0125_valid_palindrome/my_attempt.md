# 0125. Valid Palindrome - 我的原始作答

## 作答時間

2026-05-05

## 完整程式碼

```go
package main

import "fmt"

func main() {
	s := "A man, a plan, a canal: Panama"
	fmt.Println(isPalindrome(s))
}

func isPalindrome(s string) bool {
	left := 0
	right := len(s) - 1

	for right > left {
		var lb, rb byte
		for right > left {
			if isAlphaNum(s[left]) {
				lb = toLower(s[left])
				break
			}
			left++
		}
		for right > left {
			if isAlphaNum(s[right]) {
				rb = toLower(s[right])
				break
			}
			right--
		}

		if left >= right {
			return true
		}
		if lb != rb {
			return false
		}
		left++
		right--
	}

	return true
}

func isAlphaNum(ch byte) bool {
	if (ch >= 'a' && ch <= 'z') ||
		(ch >= 'A' && ch <= 'Z') ||
		(ch >= '0' && ch <= '9') {
		return true
	}
	return false
}

func toLower(ch byte) byte {
	if ch >= 'A' && ch <= 'Z' {
		ch = ch - 'A' + 'a'
	}
	return ch
}
```

## 驗證測資

這版已通過：

```go
"A man, a plan, a canal: Panama" // true
"race a car"                     // false
"a."                             // true
".a"                             // true
".,"                             // true
"0P"                             // false
""                               // true
```

## 你的思路

- 使用 `left` 和 `right` 從字串兩端往中間夾。
- 遇到非英數字元就跳過。
- 遇到英數字元後轉小寫再比較。
- 如果左右字元不同，回傳 `false`。
- 如果檢查完成都沒有不同，回傳 `true`。

## 做得好的地方

- 選對方法：two pointers。
- 沒有建立清理後的新字串，空間可以維持 `O(1)`。
- helper 拆分合理：
  - `isAlphaNum(ch byte)`
  - `toLower(ch byte)`
- 有修掉左右指標相遇時不該繼續比較未設定字元的邊界問題。

## 批改過程中的問題

### 1. 比較成功後忘記移動指標

一開始比較完 `lb` 和 `rb` 後，如果沒有移動 `left` / `right`，會一直比較同一組字元，造成死循環。

### 2. 找到字元時太早移動指標

如果取得左字元後立刻 `left++`，在 `"aa"` 這種短字串會讓右邊沒有機會取得字元。

### 3. 右邊遇到非英數時必須 `right--`

如果右邊是標點符號但沒有移動 `right`，內層迴圈會卡住。

### 4. 跳過非英數後左右相遇，應直接完成

例如：

```go
s := "a."
```

只看英數後剩下 `"a"`，應該是 `true`。因此跳過非英數後若 `left >= right`，可以直接回傳 `true`。
