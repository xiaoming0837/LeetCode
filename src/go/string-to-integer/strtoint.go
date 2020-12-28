package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	fmt.Println(myAtoi1("4193 with words"))
}

func myAtoi1(s string) int {
	s = strings.TrimSpace(s)
	if s == "" {
		return 0
	}
	var tag int
	if s[0] == 43 {
		s = s[1:]
		tag = 1
	} else if s[0] == 45 {
		s = s[1:]
		tag = -1
	} else if 48 <= s[0] && s[0] <= 57 {
		tag = 1
	} else {
		return 0
	}

	for index, str := range s {
		if str < 48 || str > 57 {
			s = s[:index]
			break
		}
	}

	var num int = 0
	for _, str := range s {
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
