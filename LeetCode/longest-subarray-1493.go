//go:build leetcode_1493

package main

import "fmt"

func countSubarrayOfTrueSize(array []int, startIdx int) int {
	result := 0
	for i := startIdx; i < len(array); i++ {
		if array[i] == 0 {
			break
		}
		result++
	}
	return result
}

func longestSubarray(nums []int) int {
	result, sequenceLen := 0, 0
	for idx := 0; idx < len(nums); {
		newSequenceLen := countSubarrayOfTrueSize(nums, idx)
		idx += newSequenceLen + 1
		result = max(result, sequenceLen+newSequenceLen)
		sequenceLen = newSequenceLen
	}
	result = max(result, sequenceLen)
	return min(result, len(nums)-1)
}

func main() {
	fmt.Println(longestSubarray([]int{0, 1, 1, 1, 0, 1, 1, 0, 1}))
}
