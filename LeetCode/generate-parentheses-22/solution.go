package main

type Parentheses struct {
	seq   string
	open  int
	close int
}

func generateParenthesis(n int) []string {
	dp := []Parentheses{{"(", 1, 0}}
	for range 2*n - 1 {
		next := make([]Parentheses, 0, len(dp)*2)
		for _, curParenth := range dp {
			if curParenth.open < n && curParenth.open >= curParenth.close {
				newParenth := Parentheses{curParenth.seq + "(", curParenth.open + 1, curParenth.close}
				next = append(next, newParenth)
			}
			if curParenth.close < n && curParenth.open > curParenth.close {
				newParenth := Parentheses{curParenth.seq + ")", curParenth.open, curParenth.close + 1}
				next = append(next, newParenth)
			}
		}
		dp = next
	}
	result := make([]string, 0, len(dp))
	for _, curParenth := range dp {
		result = append(result, curParenth.seq)
	}
	return result
}
