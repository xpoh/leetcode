package main

import (
	"fmt"
	"math"
)

// https://leetcode.com/problems/reverse-integer/

func reverse(x int) int {

	if x > math.MaxInt32 || x < math.MinInt32 {
		return 0
	}

	var minus int
	if x < 0 {
		minus = -1
	} else {
		minus = 1
	}

	d := minus * x
	s := make([]int, 0)

	for d > 0 {
		s = append(s, d%10)
		d = d / 10
	}
	fmt.Println(s)
	ans := 0
	for i, v := range s {
		ans += v * int(math.Pow(float64(10), float64(len(s)-i-1)))

	}
	ans = minus * ans

	if ans > math.MaxInt32 || ans < math.MinInt32 {
		return 0
	}

	return ans
}
