package main

import (
	"fmt"
	"math"
)

func isPalindrome(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	reversedX := 0

	for x > reversedX {
		reversedX = reversedX*10 + x%10
		x = int(math.Floor(float64(x)) / 10)
	}

	return x == reversedX || x == int(math.Floor(float64(reversedX)/10))
}

func main() {
	fmt.Printf("%t\n", isPalindrome(121))
	fmt.Printf("%t\n", isPalindrome(-121))
	fmt.Printf("%t\n", isPalindrome(10))
}
