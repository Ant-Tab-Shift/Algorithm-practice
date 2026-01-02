package main

import (
	"fmt"
)

func numSquares(n int) int {
	dp := make([]int, n+1)

	for num := 1; num*num <= n; num++ {
		square := num * num

		for sumSquare := square; sumSquare <= n; sumSquare++ {
			newCount := dp[sumSquare-square] + 1

			if dp[sumSquare] == 0 {
				dp[sumSquare] = newCount
			}
			dp[sumSquare] = min(dp[sumSquare], newCount)
		}
	}

	return dp[n]
}

func main() {
	fmt.Println(numSquares(13) == 2)
	fmt.Println(numSquares(10_000) == 1)
	fmt.Println(numSquares(19) == 3)
}
