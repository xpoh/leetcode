package main

// https://leetcode.com/problems/zigzag-conversion/

func convert(s string, numRows int) string {

	if len(s) == 1 {
		return s
	}
	if numRows == 1 {
		return s
	}

	rows := make([][]rune, numRows)
	for i := 0; i < numRows; i++ {
		rows[i] = make([]rune, len(s)/2+1)
	}

	x, y := 0, 0
	zigzag := 0
	for _, v := range s {
		if zigzag == 0 {
			rows[y][x] = v
			y += 1
			if y == numRows {
				x += 1
				y = numRows - 2
				zigzag = 1
			}
		} else {
			rows[y][x] = v
			y -= 1
			x += 1
			if y < 0 {
				y = 1
				x--
				zigzag = 0
			}
		}
	}
	var ans string

	for _, r := range rows {
		for _, b := range r {
			if b > 0 {
				ans = ans + string(b)
			}
		}
	}
	return string(ans)
}
