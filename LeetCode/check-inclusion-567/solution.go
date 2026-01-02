package main

import "fmt"

func checkInclusion(s1 string, s2 string) bool {
	if len(s2) < len(s1) {
		return false
	}

	sourceCharFreqs := make([]int, 26)
	mismatches := 0
	for idx := 0; idx < len(s1); idx++ {
		if sourceCharFreqs[s1[idx]-'a'] == 0 {
			mismatches++
		} else if sourceCharFreqs[s1[idx]-'a'] == -1 {
			mismatches--
		}
		sourceCharFreqs[s1[idx]-'a']++

		if sourceCharFreqs[s2[idx]-'a'] == 0 {
			mismatches++
		} else if sourceCharFreqs[s2[idx]-'a'] == 1 {
			mismatches--
		}
		sourceCharFreqs[s2[idx]-'a']--
	}

	for rightIdx := len(s1); rightIdx < len(s2); rightIdx++ {
		if mismatches == 0 {
			return true
		}

		if sourceCharFreqs[s2[rightIdx]-'a'] == 1 {
			mismatches--
		} else if sourceCharFreqs[s2[rightIdx]-'a'] == 0 {
			mismatches++
		}
		sourceCharFreqs[s2[rightIdx]-'a']--

		if sourceCharFreqs[s2[rightIdx-len(s1)]-'a'] == 0 {
			mismatches++
		} else if sourceCharFreqs[s2[rightIdx-len(s1)]-'a'] == -1 {
			mismatches--
		}
		sourceCharFreqs[s2[rightIdx-len(s1)]-'a']++
	}

	return mismatches == 0
}

func main() {
	fmt.Println(checkInclusion("ab", "eidbaooo"))
	fmt.Println(checkInclusion("adc", "dcda"))
	fmt.Println(checkInclusion("a", "ab"))
}
