package main

import "math"

func mathSolution(n int) int {
	sqrt := int(math.Sqrt(float64(n)))
	if sqrt*sqrt == n {
		return 1
	}

	nTemp := n
	for nTemp%4 == 0 {
		nTemp /= 4
	}
	if nTemp%8 == 7 {
		return 4
	}

	for i := 1; i*i < n; i++ {
		j := int(math.Sqrt(float64(n - i*i)))
		if i*i+j*j == n {
			return 2
		}
	}

	return 3
}
