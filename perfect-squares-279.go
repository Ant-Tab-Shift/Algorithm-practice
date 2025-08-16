//go:build leetcode_279

package main

import "fmt"

func numSquares(n int) int {
	seiveSquares := make([]int, n+1)
	for num := 1; num*num <= n; num++ {
		square := num * num
		for sumSquare := square; sumSquare <= n; sumSquare++ {
			newSumApproach := seiveSquares[sumSquare-square] + 1
			if seiveSquares[sumSquare] == 0 {
				seiveSquares[sumSquare] = newSumApproach
			} else {
				seiveSquares[sumSquare] = min(seiveSquares[sumSquare], newSumApproach)
			}
		}
	}
	return seiveSquares[n]
}

func main() {
	fmt.Println(numSquares(1) == 1)
	fmt.Println(numSquares(10_000) == 1)
	fmt.Println(numSquares(12) == 3)
}
