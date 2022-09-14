package main

import (
	"math"
)

//https://leetcode.com/problems/string-to-integer-atoi/

func myAtoi(s string) int {
	n := make([]rune, 0)
	str := "0123456789"
	m := make(map[rune]int)
	for i, v := range str {
		m[v] = i
	}

	startNum := false
	sign := 1

	for _, r := range s {
		if r == ' ' {
			if startNum {
				break
			}
			continue
		} else if r == '-' {
			if startNum {
				break
			}
			startNum = true
			sign = -1
		} else if r == '+' {
			if startNum {
				break
			}
			startNum = true
		} else if r >= '0' && r <= '9' {
			n = append(n, r)
			startNum = true
		} else if r == '.' {
			break
		} else {
			break
		}
	}
	ans := int64(0)
	d := 0
	notZero := false
	for i, r := range n {
		if r != '0' {
			notZero = true
			d = i
			break
		}
	}
	if !notZero {
		return 0
	}
	n = n[d:]
	if len(n) > 10 {
		if sign == 1 {
			return math.MaxInt32
		} else {
			return math.MinInt32
		}
	}
	if !startNum {
		return 0
	} else {
		for i, v := range n {
			ans += int64(m[v]) * int64(math.Pow(float64(10), float64(len(n)-i-1)))
		}
		ans = int64(sign) * ans
	}

	if ans > math.MaxInt32 {
		ans = math.MaxInt32
	}
	if ans < math.MinInt32 {
		ans = math.MinInt32
	}

	return int(ans)
}
