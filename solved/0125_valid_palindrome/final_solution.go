package main

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
