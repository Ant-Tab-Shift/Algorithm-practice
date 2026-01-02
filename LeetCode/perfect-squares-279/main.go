package main

import (
	"fmt"
)

func numSquares(n int) int {
	dp := make([]int, n+1)

	for sum := 1; sum <= n; sum++ {
		dp[sum] = n
		for num := 1; num*num <= sum; num++ {
			dp[sum] = min(dp[sum], dp[sum-num*num] + 1)
		}
	}

	return dp[n]
}

func main() {
	fmt.Println(numSquares(13) == 2)
	fmt.Println(numSquares(10_000) == 1)
	fmt.Println(numSquares(19) == 3)
}
