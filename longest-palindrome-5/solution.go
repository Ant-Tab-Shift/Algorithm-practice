package main

import (
	"fmt"
	"strings"
)

func manacher(s string) (resL, resR int) {
	unifiedS := "#" + strings.Join(strings.Split(s, ""), "#") + "#"
	dp := make([]int, len(unifiedS)) // distance from the end to the center of the palindrome
	l, r := 0, 0                     // previous bounds of the palindrome, which has the farthest right boundary
	for i := range len(unifiedS) {
		if r > i {
			dp[i] = min(r-i, dp[l+(r-i)])
		}
		newL, newR := i-dp[i]-1, i+dp[i]+1
		for newL >= 0 && newR < len(unifiedS) && unifiedS[newL] == unifiedS[newR] {
			dp[i]++
			newL, newR = newL-1, newR+1
		}
		if i+dp[i] > r {
			l, r = i-dp[i], i+dp[i]
			if r-l > resR-resL {
				resL, resR = l, r
			}
		}
	}
	return resL / 2, (resR + 1) / 2
}

func longestPalindrome(s string) string {
	resL, resR := manacher(s)
	return s[resL:resR]
}

func main() {
	fmt.Println(longestPalindrome("a") == "a")
	fmt.Println(longestPalindrome("aa") == "aa")
	fmt.Println(longestPalindrome("aaa") == "aaa")
	fmt.Println(longestPalindrome("aaaa") == "aaaa")
	fmt.Println(longestPalindrome("aaaaa") == "aaaaa")
	fmt.Println(longestPalindrome("acfb") == "a")
	fmt.Println(longestPalindrome("lkabac") == "aba")
}
