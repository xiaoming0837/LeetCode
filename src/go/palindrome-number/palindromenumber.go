package main

import "fmt"

func main() {
	// fmt.Println(isPalindrome(121))
	// fmt.Println(isPalindrome(-121))
	// fmt.Println(isPalindrome(10))
	// fmt.Println(isPalindrome(-101))
	fmt.Println(isPalindrome(1000030001))
}

// isPalindrome
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	str := fmt.Sprintf("%d", x)
	if len(str) == 1 {
		return true
	}
	if len(str) == 2 {
		return str[0] == str[1]
	}
	for i, j := 0, len(str)-1; i < j; i, j = i+1, j-1 {
		if str[i] != str[j] {
			return false
		}
	}
	return true
}
