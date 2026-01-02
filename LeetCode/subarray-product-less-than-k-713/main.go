package main

import "fmt"

func numSubarrayProductLessThanK(nums []int, k int) int {
	if k <= 1 {
		return 0
	}

	count := 0
	left, product := 0, 1
	for right := range len(nums) {
		product *= nums[right]

		for product >= k {
			product /= nums[left]
			left++
		}

		count += right - left + 1
	}

	return count
}

func main() {
	fmt.Println(numSubarrayProductLessThanK([]int{10, 5, 2, 6}, 100) == 8)
	fmt.Println(numSubarrayProductLessThanK([]int{1, 2, 3}, 0) == 0)
}
