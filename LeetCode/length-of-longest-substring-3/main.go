package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	maxDistance := 0
	left := -1
	lastOccurs := make(map[rune]int, 50)
	for right, char := range s {
		if prevOccur, ok := lastOccurs[char]; ok {
			left = max(left, prevOccur)
		}

		lastOccurs[char] = right
		maxDistance = max(maxDistance, right-left)
	}

	return maxDistance
}

func main() {
	fmt.Println(lengthOfLongestSubstring("") == 0)
	fmt.Println(lengthOfLongestSubstring("aab") == 2)
	fmt.Println(lengthOfLongestSubstring("bb") == 1)
}
