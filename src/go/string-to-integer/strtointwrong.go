package main

import (
	"math"
)

// Wrong
func myAtoi(s string) int {
	start, end := 0, 0
	for index, str := range s {
		if str == 43 || str == 45 || (48 <= str && str <= 57) {
			start = index
			break
		}
	}
	for index, str := range s[start+1:] {
		if str < 48 || str > 57 {
			end = index + start + 1
			break
		}
	}
	var result string
	if end == 0 {
		result = s[start:]
	} else {
		result = s[start:end]
	}
	var tag int
	if result[0] == 45 {
		tag = -1
		result = result[1:]
	} else {
		if result[0] == 43 {
			result = result[1:]
		}
		tag = 1
	}
	var num int = 0
	for _, str := range result {
		num = num*10 + int((str - '0'))
		if tag*num < math.MinInt32 {
			return math.MinInt32
		}
		if tag*num > math.MaxInt32 {
			return math.MaxInt32
		}
	}
	return int(num * tag)
}
