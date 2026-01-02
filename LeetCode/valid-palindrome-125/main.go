package main

import (
	"fmt"
	"unicode"
)

func isAlphaNum(char byte) bool {
	sym := rune(char)
	return unicode.IsLetter(sym) || unicode.IsDigit(sym)
}

func lower(char byte) rune {
	return unicode.ToLower(rune(char))
}

func isPalindrome(s string) bool {
	leftPtr, rightPtr := 0, len(s)-1
	for leftPtr < rightPtr {
		for leftPtr < rightPtr && !isAlphaNum(s[leftPtr]) {
			leftPtr++
		}
		for rightPtr > leftPtr && !isAlphaNum(s[rightPtr]) {
			rightPtr--
		}

		if lower(s[leftPtr]) != lower(s[rightPtr]) {
			return false
		}
		leftPtr++
		rightPtr--
	}

	return true
}

func main() {
	fmt.Println(isPalindrome("."))
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
}
