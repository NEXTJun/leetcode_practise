package main

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
