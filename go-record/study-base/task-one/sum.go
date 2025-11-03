package main

import (
	"fmt"
	"slices"
)

func twoSum(nums []int, target int) []int {
	for i, v := range nums {
		var current = target - v
		var idx = slices.Index(nums, current)
		if idx >= 0 && idx != i {
			return []int{i, idx}
		}
	}

	return nil
}

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
	fmt.Println(twoSum([]int{3, 2, 4}, 6))
	fmt.Println(twoSum([]int{3, 3}, 6))
}
