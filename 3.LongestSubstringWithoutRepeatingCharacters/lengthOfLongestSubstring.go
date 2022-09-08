package main

func lengthOfLongestSubstring(s string) int {
	var n, nmax int
	for k, _ := range s {
		m := make(map[byte]struct{})
		n = 0
		for i := k; i < len(s); i++ {
			if _, ok := m[s[i]]; !ok {
				m[s[i]] = struct{}{}
				n++
			} else {
				break
			}
		}
		if n > nmax {
			nmax = n
		}
	}
	if n > nmax {
		nmax = n
	}
	return nmax
}
