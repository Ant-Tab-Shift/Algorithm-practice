package main

import "fmt"

func checkSubarraySum(nums []int, k int) bool {
	memo := make(map[int]int)
	memo[0] = -1

	prefixSum := 0
	for idx, num := range nums {
		prefixSum += num

		prevIdx, ok := memo[prefixSum%k]
		if ok && idx-prevIdx >= 2 {
			return true
		}

		if !ok {
			memo[prefixSum%k] = idx
		}
	}

	return false
}

func main() {
	fmt.Println(checkSubarraySum([]int{1, 0}, 1))
	fmt.Println(checkSubarraySum([]int{1, 0}, 2))
}
