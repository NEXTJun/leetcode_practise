package main

import "sort"

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	bs := []byte(s)
	bt := []byte(t)

	sort.Slice(bs, func(i, j int) bool {
		return bs[i] < bs[j]
	})
	sort.Slice(bt, func(i, j int) bool {
		return bt[i] < bt[j]
	})

	return string(bs) == string(bt)
}
