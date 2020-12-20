package main

import "fmt"

func main() {
	//[[4,5],[2,4],[4,6],[3,4],[0,0],[1,1],[3,5],[2,2]]
	intervals := [][]int{
		{4, 5},
		{2, 4},
		{4, 6},
		{3, 4},
		{0, 0},
		{1, 1},
		{3, 5},
		{2, 2},
	}
	intervals = merge(intervals)
	fmt.Printf("after merge : %+v\n", intervals)
}

// merge intervals
func merge(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}
	// sort
	intervals = quikSortArray(intervals, 0, len(intervals)-1)
	var result [][]int
	result = append(result, intervals[0])
	fmt.Printf("befort merge : %+v\n", intervals)
	// merge
	for _, value := range intervals {
		if result[len(result)-1][1] >= value[0] && result[len(result)-1][1] < value[1] {
			result[len(result)-1][1] = value[1]
		} else if value[0] > result[len(result)-1][1] {
			result = append(result, value)
		}
	}
	return result
}

// sort intervals with frist value
func quikSortArray(intervals [][]int, left, right int) [][]int {
	if left > right {
		return intervals
	}
	// pivot
	pivot := intervals[left][0]
	i, j := left, right
	for i != j {
		// from right to left
		for i != j {
			if intervals[j][0] < pivot {
				// swap
				tmp := intervals[i]
				intervals[i] = intervals[j]
				intervals[j] = tmp
				i++
				break
			} else {
				j--
			}
		}

		// from left to right
		for i != j {
			if intervals[i][0] > pivot {
				// swap
				tmp := intervals[j]
				intervals[j] = intervals[i]
				intervals[i] = tmp
				j--
				break
			} else {
				i++
			}
		}
	}
	quikSortArray(intervals, left, i-1)
	quikSortArray(intervals, i+1, right)
	return intervals
}
