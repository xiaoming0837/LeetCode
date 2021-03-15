package main

import "math"

func main() {
	divide(1, 2)
}

func divide(dividend int, divisor int) int {

	if math.MaxInt32 < dividend || math.MaxInt32 < divisor || math.MinInt32 > dividend || math.MinInt32 > divisor {
		return math.MaxInt32
	}

	if dividend == 0 {
		return 0
	}

	f := math.Abs(float64(dividend))
	s := math.Abs(float64(divisor))
	ans := 0
	for f >= s {
		f = f - s
		ans++
	}

	if dividend < 0 {
		ans = -ans
	}
	if divisor < 0 {
		ans = -ans
	}

	return ans
}
