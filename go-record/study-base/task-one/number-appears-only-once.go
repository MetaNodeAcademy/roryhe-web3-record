package main

import "fmt"

/**
136. 只出现一次的数字：给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。
可以使用 for 循环遍历数组，结合 if 条件判断和 map 数据结构来解决，例如通过 map 记录每个元素出现的次数，然后再遍历 map 找到出现次数为1的元素。
*/
// 因x^0=x,x^x=0,x^y=y^x,x^y^z....=....z^y^x,所以但凡有相同数异或值都为0，循环找到一个不为0的值则是唯一数
func singleNumber(nums []int) int {
	result := 0
	for i := 0; i < len(nums); i++ {
		result ^= nums[i]
	}
	return result
}

func singleNumberByMap(nums []int) int {
	var resultMap map[int]int = make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if len(resultMap) == 0 {
			resultMap[nums[i]] = 1
		} else if _, key := resultMap[nums[i]]; key {
			resultMap[nums[i]] += 1
		} else {
			resultMap[nums[i]] = 1
		}
	}

	for key, value := range resultMap {
		if value == 1 {
			return key
		}
	}
	return 0
}

func main() {
	fmt.Println(singleNumber([]int{2, 2, 1}))
	fmt.Println(singleNumber([]int{4, 1, 2, 1, 2}))
	fmt.Println(singleNumber([]int{1}))

	fmt.Println(singleNumberByMap([]int{2, 2, 1}))
	fmt.Println(singleNumberByMap([]int{4, 1, 2, 1, 2}))
	fmt.Println(singleNumberByMap([]int{1}))
}
