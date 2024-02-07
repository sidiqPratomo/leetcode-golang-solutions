package medianoftwosortedarrays

import "math"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		return findMedianSortedArrays(nums2, nums1)
	}
	m := len(nums1)
	n := len(nums2)
	low := 0
	high := m
	for low <= high {
		px := (low + high) / 2
		py := ((m + n + 1) / 2) - px

		var maxLeft1, minRight1, maxLeft2, minRight2 int
		if px == 0 {
			maxLeft1 = math.MinInt64
		} else {
			maxLeft1 = nums1[px-1]
		}
		if px == m {
			minRight1 = math.MaxInt64
		} else {
			minRight1 = nums1[px]
		}

		if py == 0 {
			maxLeft2 = math.MinInt64
		} else {
			maxLeft2 = nums2[py-1]
		}
		if py == n {
			minRight2 = math.MaxInt64
		} else {
			minRight2 = nums2[py]
		}

		if maxLeft1 <= minRight2 && maxLeft2 <= minRight1 {
			if (m+n)%2 == 0 {
				return float64(math.Max(float64(maxLeft1), float64(maxLeft2))+math.Min(float64(minRight1), float64(minRight2))) / 2
			} else {
				return float64(math.Max(float64(maxLeft1), float64(maxLeft2)))
			}
		} else if maxLeft1 > minRight2 {
			high = px - 1
		} else {
			low = px + 1
		}
	}
	return 0
}