package main

import "fmt"

func main() {
	// fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}))
	// fmt.Println(longestCommonPrefix([]string{"dog", "racecar", "car"}))
	// fmt.Println(longestCommonPrefix([]string{"a"}))
	// fmt.Println(longestCommonPrefix([]string{"", ""}))
	// fmt.Println(longestCommonPrefix([]string{}))
	fmt.Println(longestCommonPrefix([]string{"ab", "a"}))
}

func longestCommonPrefix(strs []string) string {
	l := len(strs)
	result := ""

	for i := 0; i < 200 && l != 0; i++ {
		for j := 1; j < l; j++ {

			if strs[j] == "" || strs[j-1] == "" {
				return result
			}

			if i >= len(strs[j]) || i >= len(strs[j-1]) {
				return result
			}

			if (strs[j-1])[i] != (strs[j])[i] {
				return result
			}
		}
		if i < len(strs[0]) {
			result += string((strs[0])[i])
		}
	}
	return result
}
