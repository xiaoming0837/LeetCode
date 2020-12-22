package main

import "fmt"

func main() {
	fmt.Printf("%.5f\n", findMedianSortedArrays([]int{1, 2}, []int{3, 4}))
}

// Find Median Sorted Arrays
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	arr := append(nums1, nums2...)
	fmt.Printf("%+v\n", arr)
	fmt.Printf("%+v\n", quickSorted(arr, 0, len(arr)-1))
	len := len(arr)
	if len%2 == 0 {
		return float64(arr[(len/2)-1]+arr[len/2]) / 2
	}
	return float64(arr[len/2])

}

// Quick Sorted
func quickSorted(arr []int, left, right int) []int {
	if left > right {
		return arr
	}
	// povit
	povit := arr[left]
	i, j := left, right
	for i != j {
		// from right to left
		for i != j {
			if arr[j] < povit {
				//swap
				tmp := arr[i]
				arr[i] = arr[j]
				arr[j] = tmp
				i++
				break
			} else {
				j--
			}
		}
		// from left to right
		for i != j {
			if arr[i] > povit {
				//swap
				tmp := arr[i]
				arr[i] = arr[j]
				arr[j] = tmp
				j--
				break
			} else {
				i++
			}
		}
	}
	// left
	quickSorted(arr, left, i-1)
	// right
	quickSorted(arr, i+1, right)
	return arr
}
