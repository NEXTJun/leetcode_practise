package main

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
