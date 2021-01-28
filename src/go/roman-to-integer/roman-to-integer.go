package main

import "fmt"

func main() {
	fmt.Println(romanToInt("III"))
	fmt.Println(romanToInt("IV"))
	fmt.Println(romanToInt("IX"))
	fmt.Println(romanToInt("LVIII"))
	fmt.Println(romanToInt("MCMXCIV"))
}

func romanToInt(s string) int {
	r := map[string]int{"M": 0, "D": 1, "C": 2, "L": 3, "X": 4, "V": 5, "I": 6}
	v := []int{1000, 500, 100, 50, 10, 5, 1}
	l := len(s)
	result := v[r[string(s[l-1])]]
	for i := 0; i < l-1; i++ {
		if r[string(s[i])] <= r[string(s[i+1])] {
			result += v[r[string(s[i])]]
		} else {
			result -= v[r[string(s[i])]]
		}
	}
	return result
}
