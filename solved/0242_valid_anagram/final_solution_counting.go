package main

func isAnagramCounting(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	var counts [26]int

	for _, ch := range s {
		counts[ch-'a']++
	}

	for _, ch := range t {
		counts[ch-'a']--
	}

	for _, count := range counts {
		if count != 0 {
			return false
		}
	}

	return true
}
