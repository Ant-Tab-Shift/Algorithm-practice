package main

import "fmt"

func twoSum(nums []int, target int) []int {
	numsTable := make(map[int]int, len(nums))
	for idx, num := range nums {
		if pairIdx, ok := numsTable[target-num]; ok {
			return []int{pairIdx, idx}
		}
		numsTable[num] = idx
	}
	return nil
}

func main() {
	fmt.Println(twoSum([]int{-1, -3, 10}, 9))   // 0, 2
	fmt.Println(twoSum([]int{-10, -100}, -110)) // 0, 1
}
