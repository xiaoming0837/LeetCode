package main

import (
	"fmt"
)

func main() {
	// test()
	fmt.Println(isValid("()"))
	fmt.Println(isValid("()[]{}"))
	fmt.Println(isValid("(]"))
	fmt.Println(isValid("([)]"))
	fmt.Println(isValid("{[]}"))
}

func test() {
	arr := make([]rune, 0, 0)
	arr = append(arr, rune('['))
	arr = append(arr, rune(']'))
	arr = arr[0:2]
	fmt.Println("arr")
	fmt.Println(arr)
}

func isValid(s string) bool {
	stack := make([]rune, 0, 0)
	for i := 0; i < len(s); i++ {
		if len(stack) == 0 {
			stack = add(stack, rune(s[i]))
		} else {
			if match(get(stack), rune(s[i])) {
				stack = del(stack)
			} else {
				stack = add(stack, rune(s[i]))
			}
		}
	}
	return len(stack) == 0

}

func match(r1, r2 rune) bool {
	return (r1 == '[' && r2 == ']') || (r1 == '(' && r2 == ')') || (r1 == '{' && r2 == '}')
}

func add(s []rune, n rune) []rune {
	s = append(s, n)
	return s
}

func del(s []rune) []rune {
	return s[:len(s)-1]

}
func get(s []rune) rune {
	return s[len(s)-1]
}
