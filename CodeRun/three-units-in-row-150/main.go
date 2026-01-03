package main

import (
	"fmt"
)

const MAX_UNITS_IN_ROW = 3

// Dynamic Programming: Bottom-Up with tabulation
func bottomUpSolution(N int) int {
	if N == 0 {
		return 1
	}

	// the index position define how many units in a row in the end of sequence
	table := make([]int, MAX_UNITS_IN_ROW)
	table[0] = 1

	for range N + 1 {
		sum := 0
		for idx := MAX_UNITS_IN_ROW - 1; idx > 0; idx-- {
			sum += table[idx]
			table[idx] = table[idx-1]
		}

		table[0] += sum
	}

	return table[0]
}

func main() {
	var N int
	fmt.Scan(&N)

	fmt.Println(bottomUpSolution(N))
}
