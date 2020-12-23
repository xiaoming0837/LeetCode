package main

import "fmt"

func main() {
	longestPalindrome("abcdcba")
}

func longestPalindrome(s string) string {
	if s == "" {
		return ""
	}
	// 保存起始位置
	var ran []int = make([]int, 2)
	for i := 0; i < len(s); i++ {
		//把回文看成中间的部分全是同意字符，左右部分想对称
		fmt.Println("-------------------------------------------------------")
		fmt.Printf("%+v\n", ran)
		i = findLongest(s, i, ran[:])
		fmt.Printf("%s\n", s[ran[0]:ran[1]+1])
		fmt.Println("-------------------------------------------------------")
	}
	return s[ran[0] : ran[1]+1]
}

func findLongest(s string, low int, ran []int) int {
	// 查找中间部分
	high := low
	fmt.Printf("high: %d==== low %d \n", high, low)
	for high < len(s)-1 && s[high+1] == s[low] {
		high++
	}

	// 定位中间部分的最后一个字符
	ans := high
	// 从中间向左右扩散
	for low > 0 && high < len(s)-1 && s[low-1] == s[high+1] {
		low--
		high++
	}
	// 记录最大长度
	if high-low > ran[1]-ran[0] {
		fmt.Println("复制")
		ran[0] = low
		ran[1] = high
	}
	fmt.Printf("====%+v====\n", ran)
	return ans
}
