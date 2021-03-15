package main

import "fmt"

func main() {
	nums := []int{1, 1, 2}
	removeDuplicates(nums)

}

func removeDuplicates(nums []int) int {
	m := make(map[int]bool)
	k := 0
	for len(nums) > k {
		if m[nums[k]] != false {
			nums = remove(nums, k)
		} else {
			m[nums[k]] = true
			k++
		}

	}
	fmt.Println(nums)
	return len(nums)
}

func remove(nums []int, k int) []int {
	fmt.Println(k)
	// fmt.Println(nums[:k])
	// fmt.Println(nums[k+2:])
	nums = append(nums[:k], nums[k+2:]...)

	return nums
}
