package main

import "fmt"

func main() {
	fmt.Println(letterCombinations("234"))
}

func letterCombinations(digits string) []string {
	dig := map[string][]string{
		"1": {""},
		"2": {"a", "b", "c"},
		"3": {"d", "e", "f"},
		"4": {"g", "h", "i"},
		"5": {"j", "k", "l"},
		"6": {"m", "n", "o"},
		"7": {"p", "q", "r", "s"},
		"8": {"t", "u", "v"},
		"9": {"w", "x", "y", "z"},
	}
	result := make([]string, 0, 0)
	if digits == "" {
		return result
	}

	result = append(result, dig[string(digits[0])]...)
	// tmp := make([]string, 0)
	for i := 1; i < len(digits); i++ {
		tmp := descartes(result, dig[string(digits[i])])
		result = result[:0]
		result = append(result, tmp...)
	}

	return result
}

func descartes(s1, s2 []string) []string {
	result := make([]string, 0, 0)
	for _, v1 := range s1 {
		for _, v2 := range s2 {
			result = append(result, v1+v2)
		}
	}
	return result
}
