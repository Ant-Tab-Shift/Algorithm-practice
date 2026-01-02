package main

import "fmt"

func checkSubarraySum(nums []int, k int) bool {
	memo := make(map[int]int)
	memo[0] = -1

	sum := 0
	for idx, num := range nums {
		sum += num

		prevIdx, ok := memo[sum%k]
		if ok && idx-prevIdx > 1 {
			return true
		}

		if !ok {
			memo[sum%k] = idx
		}
	}

	return false
}

func main() {
	fmt.Println(checkSubarraySum([]int{1, 0}, 1))
	fmt.Println(checkSubarraySum([]int{1, 0}, 2))
}
