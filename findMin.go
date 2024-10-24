package main

// https://leetcode.com/problems/find-minimum-in-rotated-sorted-array/
func findMin(nums []int) int {
	start, end := 0, len(nums)-1

	for start < end {
		mid := (start + end) / 2

		if nums[mid] > nums[end] {
			start = mid + 1
		} else {
			end = mid
		}
	}

	return nums[start]
}
