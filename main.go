package main

import "fmt"

func rob(nums []int) int {
	// TODO: 定義「考慮到目前房子時的最佳金額」，逐間比較拿與不拿。
	// 提醒：拿目前這間時，不能接上相鄰那間的結果。
	nl := len(nums)
	if nl == 0 {
		return 0
	} else if nl == 1 {
		return nums[0]
	}

	F := make([]int, nl)
	for i := 0; i < nl; i++ {
		if i == 0 {
			F[i] = nums[i]
			continue
		} else if i == 1 {
			F[i] = max(nums[i-1], nums[i])
			continue
		}
		F[i] = max(F[i-1], F[i-2]+nums[i])
	}
	return F[nl-1]
}

func max(n1, n2 int) int {
	if n1 > n2 {
		return n1
	}
	return n2
}

func main() {
	// 預期：12、4、5
	fmt.Println(rob([]int{2, 7, 9, 3, 1}))
	fmt.Println(rob([]int{1, 2, 3, 1}))
	fmt.Println(rob([]int{5}))
}
