# 0020. Valid Parentheses - 我的原始作答

## 作答時間

2026-05-03

## 完整程式碼

```go
package main

import "fmt"

func main() {
	s := "([])"
	fmt.Println(isValid(s))
}

func isValid(s string) bool {
	stack := []rune{}
	var res bool

	for i, ch := range s {
		if ch == '{' || ch == '[' || ch == '(' {
			var w rune
			switch ch {
			case '(':
				w = ')'
			case '[':
				w = ']'
			case '{':
				w = '}'
			}
			stack = append(stack, w)
		} else if ch == ')' || ch == ']' || ch == '}' {
			if len(stack) > 0 {
				if stack[len(stack)-1] == ch {
					stack = stack[:len(stack)-1]
				}
			} else {
				break
			}
			if i == len(stack) {
				res = true
			}
		} else {
			break
		}

	}

	return res
}
```

## 你的思路

- 使用 stack 處理括號配對。
- 遇到左括號時，不直接存左括號，而是存「預期遇到的右括號」。
- 遇到右括號時，檢查 stack 最上層是否等於目前字元。

## 做得好的地方

- 選對資料結構：這題核心就是 stack。
- 「stack 裡存預期的右括號」是很好的寫法，比存左括號後再轉換更直接。
- 已經有處理 stack pop 的概念：

```go
stack = stack[:len(stack)-1]
```

## 需要修正的地方

### 1. `res` 狀態不需要

這題最後只需要判斷：

```go
return len(stack) == 0
```

過程中只要遇到錯誤就立刻 `return false`。

### 2. `i == len(stack)` 判斷沒有意義

```go
if i == len(stack) {
	res = true
}
```

`i` 是字串的 byte index，`len(stack)` 是 stack 長度，兩者不是同一種概念，不能拿來判斷括號是否配對完成。

### 3. 右括號不匹配時沒有立刻失敗

如果目前右括號和 stack top 不一樣，應該直接：

```go
return false
```

例如：

```go
s := "([)]"
```

這種順序錯誤的括號必須回傳 `false`。

### 4. stack 為空又遇到右括號時，應該直接失敗

例如：

```go
s := ")"
```

沒有任何左括號，卻先遇到右括號，應直接回傳 `false`，不是 `break`。
