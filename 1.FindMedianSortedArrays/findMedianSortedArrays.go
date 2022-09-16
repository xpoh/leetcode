package main

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	ans := make([]int, 0)
	i1, i2 := 0, 0

	for i1 < len(nums1) || i2 < len(nums2) {

		if i1 < len(nums1) && i2 < len(nums2) {
			if nums1[i1] < nums2[i2] {
				ans = append(ans, nums1[i1])
				i1++
			} else {
				ans = append(ans, nums2[i2])
				i2++
			}
		} else {
			if i1 < len(nums1) {
				ans = append(ans, nums1[i1])
				i1++
			} else {
				ans = append(ans, nums2[i2])
				i2++
			}
		}
	}
	n := len(ans)
	f := float64(0)
	if n%2 == 0 {
		f = float64(ans[n/2-1]+ans[n/2]) * 0.5
	} else {
		f = float64(ans[n/2])
	}
	return f
}
