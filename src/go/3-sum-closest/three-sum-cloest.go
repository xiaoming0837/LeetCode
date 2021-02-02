package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(threeSumClosest([]int{-1, 2, 1, -4}, 1))

}

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	l := len(nums)
	result := target
	for first := 0; first < l; first++ {
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}
		third := l - 1
		for second := first + 1; second < third; {
			if second > first+1 && nums[second] == nums[second-1] {
				second++
				continue
			}
			val := nums[first] + nums[second] + nums[third]
			if val == target {
				return target
			}
			if val < target {
				second++
			} else {
				third--
			}
			if result == target || abs(val-target) < abs(result-target) {
				fmt.Printf("%d %d %d \n", nums[first], nums[second], nums[third])
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
