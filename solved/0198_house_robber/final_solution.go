package main

func rob(nums []int) int {
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
