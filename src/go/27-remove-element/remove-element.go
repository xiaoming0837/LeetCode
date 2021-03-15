package main

func main() {

}

func removeElement(nums []int, val int) int {
	k := 0
	for len(nums) > k {
		if nums[k] == val {
			tmp := nums[len(nums)-1]
			nums[k] = tmp
			nums = nums[:len(nums)-1]
		} else {
			k++
		}
	}
	return len(nums)
}
