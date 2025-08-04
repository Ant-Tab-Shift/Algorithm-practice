package main

import "strings"

func isAlnum(sym byte) bool {
	return sym <= 'z' && sym >= 'a' || sym <= 'Z' && sym >= 'A' || sym <= '9' && sym >= '0'
}

func lower(sym byte) string {
	return strings.ToLower(string(sym))
}

func isPalindrome(s string) bool {
	leftPtr := 0
	rightPtr := len(s) - 1
	for leftPtr < rightPtr {
		if !isAlnum(s[leftPtr]) {
			leftPtr++
		} else if !isAlnum(s[rightPtr]) {
			rightPtr--
		} else if lower(s[leftPtr]) != lower(s[rightPtr]) {
			return false
		} else {
			leftPtr++
			rightPtr--
		}
	}
	return true
}
