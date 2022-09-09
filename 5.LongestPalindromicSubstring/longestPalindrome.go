package main

func isPalindrom(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}

func longestPalindrome(s string) string {
	var ans string
	n := 0
	if len(s) == 1 {
		return s
	}
	for i := 0; i <= len(s); i++ {
		for j := i; j <= len(s); j++ {
			if isPalindrom(s[i:j]) && len(s[i:j]) > n {
				n = len(s[i:j])
				ans = s[i:j]
			}
		}
	}
	return ans
}
