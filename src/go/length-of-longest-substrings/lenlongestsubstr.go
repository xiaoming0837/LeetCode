package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	// fmt.Println(lengthOfLongestSubstring("bbbbb"))
	// fmt.Println(lengthOfLongestSubstring("pwwkew"))
	// fmt.Println(lengthOfLongestSubstring(""))
	// fmt.Println(lengthOfLongestSubstring("au"))
	// fmt.Println(lengthOfLongestSubstring("bwf"))
}

// lengthOfLongestSubstring
func lengthOfLongestSubstring(s string) int {
	if s == "" {
		return 0
	}
	tmp := 1
	for i := 2; i <= len(s); i++ {
		for j := 0; j+i <= len(s); j++ {
			val := s[j : j+i]
			fmt.Printf("---%s\n", val)
			if isUniqueString(val) {
				if len(val) > tmp {
					tmp++
				}
				break
			}
		}

		fmt.Printf("%d===%d\n", tmp, i)
		if tmp != i {
			break
		}
	}

	return tmp
}

// isUniqueString
func isUniqueString(s string) bool {
	for _, v := range s {
		if strings.Count(s, string(v)) > 1 {
			return false
		}
	}
	return true
}
