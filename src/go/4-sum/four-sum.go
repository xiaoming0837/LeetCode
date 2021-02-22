package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(fourSum([]int{1, 0, -1, 0, -2, 2}, 0))
	fmt.Println(fourSum([]int{}, 0))
}

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	l := len(nums)
	ans := make([][]int, 0, 0)
	for first := 0; first < l; first++ {
		if first > 0 && nums[first-1] == nums[first] {
			continue
		}

		for second := first + 1; second < l; second++ {
			if second > first+1 && nums[second-1] == nums[second] {
				continue
			}
			fourth := l - 1
			for third := second + 1; third < fourth; {
				if third > second+1 && nums[third-1] == nums[third] {
					third++
					continue
				}
				if nums[first]+nums[second]+nums[third]+nums[fourth] < target {
					third++
				} else if nums[first]+nums[second]+nums[third]+nums[fourth] > target {
					fourth--
				} else {
					ans = append(ans, []int{nums[first], nums[second], nums[third], nums[fourth]})
					third++
					fourth--
				}

			}
		}
	}
	return ans
}
