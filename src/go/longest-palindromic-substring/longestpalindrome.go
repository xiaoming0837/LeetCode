package main

import "fmt"

func main() {
	fmt.Printf("%s\n", longestPalindrome("babad"))
	fmt.Printf("%s\n", longestPalindrome("cbbd"))
	fmt.Printf("%s\n", longestPalindrome("a"))
	fmt.Printf("%s\n", longestPalindrome("ac"))
	fmt.Printf("%s\n", longestPalindrome("addddadadasfgg1123"))
}

// longestPalindrome DP palindrome
func longestPalindrome(s string) string {
	n := len(s)
	ans := ""
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}

	for l := 0; l < n; l++ {
		for i := 0; l+i < n; i++ {
			j := l + i
			if l == 0 {
				dp[i][j] = 1
			} else if l == 1 {
				if s[i] == s[j] {
					dp[i][j] = 1
				}
			} else {
				if s[i] == s[j] {
					dp[i][j] = dp[i+1][j-1]
				}
			}
			if dp[i][j] > 0 && l+1 > len(ans) {
				ans = s[i : i+l+1]
			}
		}
	}
	return ans
}
