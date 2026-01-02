package main

import "fmt"

func shortestToChar(s string, c byte) []int {
	answer := make([]int, len(s))
	prevOccur := -len(s)
	for idx := range len(s) {
		if s[idx] == c {
			prevOccur = idx
		}
		answer[idx] = idx - prevOccur
	}
	
	for idx := len(s) - 2; idx >= 0; idx-- {
		answer[idx] = min(answer[idx], answer[idx+1]+1)
	}

	return answer
}

func main() {
	fmt.Println(shortestToChar("a", 'a'))     // [0]
	fmt.Println(shortestToChar("ba", 'a'))    // [1 0]
	fmt.Println(shortestToChar("abbab", 'a')) // [0 1 1 0 1]
}
