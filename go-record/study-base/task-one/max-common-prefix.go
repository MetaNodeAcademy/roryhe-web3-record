package main

import "fmt"

func maxComPrefix(strs []string) string {
	/**
	先找最短，再根据最短的string中的每一个char有序匹配其他string
	*/
	if len(strs) == 0 {
		return ""
	}

	var minStr = strs[0]
	for _, str := range strs {
		minStr = min(minStr, str)
	}

	for i := 0; i < len(minStr); i++ {
		current := minStr[i]
		for _, child := range strs {
			if child[i] != current {
				return minStr[:i]
			}
		}
	}

	return minStr
}

func main() {
	fmt.Println("first", maxComPrefix([]string{"dog", "racecar", "car"}))
	fmt.Println("second", maxComPrefix([]string{"flower", "flow", "flight"}))
}
