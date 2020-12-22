package main

import "fmt"

func main() {
	// fmt.Printf("%s\n", longestPalindrome("babad"))
	// fmt.Printf("%s\n", longestPalindrome("cbbd"))
	// fmt.Printf("%s\n", longestPalindrome("a"))
	// fmt.Printf("%s\n", longestPalindrome("ac"))
	fmt.Printf("%s\n", longestPalindrome("addddadadasfgg1123"))
}

// longestPalindrome
func longestPalindrome(s string) string {
	return palindromes(s, 0, len(s)-1)
}

// get palindromes
func palindromes(s string, left, right int) string {
	if s[left] == s[right] {
		s = s[left:right]
		if left-right <= 2 {
			return
		}
	}
	if left-right == 1 {
		return s[lef]
	}

	// left
	leftVal := palindromes(s, left, right-1)
	rightVal := palindromes(s, left+1, right)
	if len(leftVal) < len(rightVal) {
		return rightVal
	}
	return leftVal
}
