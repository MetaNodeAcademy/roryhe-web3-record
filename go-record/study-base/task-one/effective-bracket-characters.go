package main

import "fmt"

func isValid(s string) bool {
	if len(s) <= 1 || len(s)%2 != 0 {
		return false
	}

	var stack = make([]rune, 0)

	bracketMap := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	for i := 0; i < len(s); i++ {
		var char = s[i]
		if left, right := bracketMap[rune(char)]; right {
			if len(stack) == 0 || stack[len(stack)-1] != left {
				return false
			}

			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, rune(char))
		}
	}

	return len(stack) == 0
}

func main() {
	var inputs []string = []string{"()", "()[]{}", "(]", "([])", "([)]"}
	var results []bool = make([]bool, len(inputs))
	for i := 0; i < len(inputs); i++ {
		results[i] = isValid(inputs[i])
	}

	for j := range results {
		fmt.Println(results[j])
	}
}
