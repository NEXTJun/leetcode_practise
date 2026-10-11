# 0020. Valid Parentheses

## 題目

給定只包含 `(`、`)`、`[`、`]`、`{`、`}` 的字串，判斷括號是否有效。有效代表每個左括號都要被相同類型的右括號關閉，而且順序必須正確。

## 最終作答

```go
func isValid(s string) bool {
	stack := []rune{}

	for _, ch := range s {
		if ch == '(' {
			stack = append(stack, ')')
		} else if ch == '[' {
			stack = append(stack, ']')
		} else if ch == '{' {
			stack = append(stack, '}')
		} else {
			if len(stack) == 0 {
				return false
			}
			if stack[len(stack)-1] != ch {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}
```

## 解法重點

- 用 stack 處理「最後出現的左括號，必須最先被關閉」。
- 遇到左括號時，直接 push 對應的右括號。
- 遇到右括號時，檢查 stack top 是否等於目前字元。
- 如果 stack 是空的卻遇到右括號，代表沒有對應的左括號，回傳 `false`。
- 如果最後 stack 還有東西，代表有左括號沒被關閉，回傳 `false`。

## 從原始作答改進的地方

- 移除 `res`，改成錯誤時直接 `return false`。
- 移除錯誤判斷 `i == len(stack)`。
- 右括號不匹配時立刻回傳 `false`。
- stack 空時遇到右括號立刻回傳 `false`。
- 最後用 `return len(stack) == 0` 判斷是否全部配對完成。

## 複雜度

- 時間複雜度：`O(n)`
- 空間複雜度：`O(n)`

## 面試口述版

我會用 stack 來處理括號順序。因為括號必須符合後進先出，最後打開的左括號必須最先被關閉。走訪字串時，如果遇到左括號，我直接把它對應的右括號放進 stack。遇到右括號時，檢查 stack 是否為空，以及 stack top 是否等於目前字元；只要不符合就回傳 false。最後如果 stack 是空的，代表所有括號都正確配對，否則代表還有左括號沒有被關閉。時間複雜度是 `O(n)`，空間複雜度是 `O(n)`。
