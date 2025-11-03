package main

import (
	"fmt"
	"sort"
)

func merge(intervals [][]int) [][]int {
	var result [][]int

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	result = append(result, []int{intervals[0][0], intervals[0][1]})

	for i, interval := range intervals {
		lastLen := len(result) - 1
		lastItem := result[lastLen]

		if interval[0] <= lastItem[1] {
			if interval[1] > lastItem[1] {
				lastItem[1] = intervals[i][1]
			}
		} else {
			result = append(result, []int{interval[0], interval[1]})
		}
	}

	return result
}

func main() {
	fmt.Println(merge([][]int{[]int{1, 3}, []int{2, 6}, []int{8, 10}, []int{15, 18}}))
	fmt.Println(merge([][]int{[]int{4, 7}, []int{1, 4}}))
}
