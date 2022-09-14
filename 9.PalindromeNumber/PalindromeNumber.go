package main

// https://leetcode.com/problems/palindrome-number/
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	n := make([]int, 0)
	for x > 0 {
		n = append(n, x%10)
		x = x / 10
	}
	for i := 0; i < len(n)/2; i++ {
		if n[i] != n[len(n)-1-i] {
			return false
		}
	}
	return true
}
