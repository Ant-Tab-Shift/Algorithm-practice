package main

// Dynamic Programming: Top-Down with memoization
func topDownSolution(N int) int {
	memo := make(map[[2]int]int, N)

	var helper func(int, int) int
	helper = func(n, unitsInRow int) int {
		if n == 0 {
			return 1
		}

		ans, ok := memo[[2]int{n, unitsInRow}]
		if ok {
			return ans
		}

		ans = helper(n-1, 0)
		if unitsInRow < MAX_UNITS_IN_ROW-1 {
			ans += helper(n-1, unitsInRow+1)
		}

		memo[[2]int{n, unitsInRow}] = ans
		return ans
	}

	return helper(N, 0)
}
