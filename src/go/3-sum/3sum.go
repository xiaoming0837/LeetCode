package main

import "fmt"

func main() {
	// print(threeSum([]int{-1, 0, 1, 2, -1, -4}))
	// print(threeSum([]int{}))
	// print(threeSum([]int{0, 0, 0, 0}))
}

func threeSum(nums []int) [][]int {
	// init
	ans := make([][][]bool, len(nums))
	for i := range ans {
		ans[i] = make([][]bool, len(nums))
		for j := range ans[i] {
			ans[i][j] = make([]bool, len(nums))
		}
	}

	for i := 0; i < len(nums)-2; i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			for k := j + 1; k < len(nums); k++ {
				if isZero := nums[i] + nums[j] + nums[k]; isZero == 0 {
					ans[i][j][k] = true
				}
			}
		}
	}
	res := [][]int{}
	for i, v := range ans {
		for j, v2 := range v {
			for k, v3 := range v2 {
				if v3 && contain(res, sort([]int{nums[i], nums[j], nums[k]})) {
					res = append(res, sort([]int{nums[i], nums[j], nums[k]}))
				}
			}
		}
	}
	return res
}

func contain(res [][]int, a []int) bool {
	for _, val := range res {
		if val[0] == a[0] && val[1] == a[1] && val[2] == a[2] {
			return false
		}
	}
	return true
}

func sort(a []int) []int {
	if a[0] > a[1] {
		a[0], a[1] = a[1], a[0]
	}
	if a[0] > a[2] {
		a[0], a[2] = a[2], a[0]
	}
	if a[1] > a[2] {
		a[1], a[2] = a[2], a[1]
	}
	return a
}

func print(val [][]int) {
	fmt.Println(val)
}
