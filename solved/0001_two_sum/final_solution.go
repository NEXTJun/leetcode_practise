package main

func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int)

	for i, val := range nums {
		if j, ok := numMap[target-val]; ok {
			return []int{j, i}
		}
		numMap[val] = i
	}

	return nil
}
