package main

import "fmt"

func main() {
	s := "123456789"
	fmt.Println(s[0:1])
	fmt.Println(s[0:9])
	fmt.Println(longestPalindrome("abcdcba"))
}

// longestPalindrome 最长回文子串 暴力求解法
func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s[0:1]
	}
	// left, right := 0
	result := s[0:1]
	for i := 0; i < len(s)-1; i++ {
		for j := i + 1; j < len(s); j++ {
			if isPalindrome(s, i, 6) {
				if len(result) < len(s[i:j+1]) {
					result = s[i : j+1]
				}
			}
		}
	}
	return result
}

// isPalindrome
func isPalindrome(s string, left, right int) bool {
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}
