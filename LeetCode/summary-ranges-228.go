//go:build leetcode_228

package main

import "fmt"

func rangeToString(start, end int) string {
	if start != end {
		return fmt.Sprintf("%d->%d", start, end)
	}
	return fmt.Sprintf("%d", end)
}

func summaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return []string{}
	}

	result := make([]string, 0)
	start, end := 0, 1
	for ; end < len(nums); end++ {
		if nums[end-1]+1 != nums[end] {
			result = append(result, rangeToString(nums[start], nums[end-1]))
			start = end
		}
	}
	return append(result, rangeToString(nums[start], nums[end-1]))
}

func main() {
	inputNums := []int{0, 2, 3, 4, 6, 8, 9}
	fmt.Println(summaryRanges(inputNums))
}
