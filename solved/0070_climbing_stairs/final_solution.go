package main

func climbStairs(n int) int {
	if n <= 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}

	prev2 := 1
	prev1 := 2

	for i := 3; i <= n; i++ {
		current := prev1 + prev2
		prev2 = prev1
		prev1 = current
	}

	return prev1
}

func climbStairsMemo(n int) int {
	memo := make(map[int]int)

	var dfs func(int) int
	dfs = func(step int) int {
		if step == 1 {
			return 1
		}
		if step == 2 {
			return 2
		}
		if val, ok := memo[step]; ok {
			return val
		}

		ways := dfs(step-1) + dfs(step-2)
		memo[step] = ways
		return ways
	}

	return dfs(n)
}
