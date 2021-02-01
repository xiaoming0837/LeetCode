package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(threeSumCloest([]int{-1, 0, 1, 2, -1, -4}))

}

func threeSumCloest(nums []int, target int) int {
	sort.Ints(nums)
	l := len(nums)
	result := target
	for first := 0; first < l; first++ {
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}
		third := l - 1
		tmp := target - nums[first]
		for second := first + 1; second < third; {
			if second > first+1 && nums[second] == nums[second-1] {
				second++
				continue
			}
			if nums[second]+nums[third] == tmp {
				return target
			}
			val := abs(nums[second] + nums[third] + nums[first] - target)
			if val < result {
				result = val
			}
		}
	}
	return result
}

func abs(val int) int {
	if val < 0 {
		return -1 * val
	}
	return val
}
