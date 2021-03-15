package main

func main() {
	// "mississippi"
	// "pi"
	strStr("mississippi", "pi")
}

func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	if len(needle) > len(haystack) {
		return -1
	}
	ans := -1
	for i, j := 0, 0; i < len(haystack); {
		if haystack[i] == needle[j] {
			if ans == -1 {
				ans = i
			}
			j++
			i++
			if j > len(needle)-1 {
				return ans
			}
		} else {
			if ans != -1 {
				i = ans + 1
			} else {
				i++
			}
			j = 0
			ans = -1
		}
	}
	return -1
}
