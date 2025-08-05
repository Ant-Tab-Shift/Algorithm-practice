package main

import (
	"math"
)

type Parantheses struct {
	seq   string
	open  int
	close int
}

func generateParenthesis(n int) []string {
	result := make([]Parantheses, int(math.Pow(2.0, float64(n))))
	result[0] = Parantheses{"(", 1, 0}
}
