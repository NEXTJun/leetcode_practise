package main

func maxSubArray(nums []int) int {
	currentSum := nums[0]
	maxSum := nums[0]

	for i := 1; i < len(nums); i++ {
		num := nums[i]

		if currentSum < 0 {
			currentSum = num
		} else {
			currentSum += num
		}

		if currentSum > maxSum {
			maxSum = currentSum
		}
	}

	return maxSum
}
