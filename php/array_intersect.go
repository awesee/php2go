package php

// ArrayIntersect — Computes the intersection of arrays
func ArrayIntersect(nums1, nums2 []int) []int {

	res := make([]int, 0, len(nums1))
	nc := make(map[int]int)

	for _, n := range nums1 {
		nc[n]++
	}

	for _, n := range nums2 {
		if nc[n] > 0 {
			res = append(res, n)
			nc[n]--
		}
	}

	return res
}
