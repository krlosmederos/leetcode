package main

import "math"

func minEatingSpeed(piles []int, h int) int {

	maxPiles := 0
	for _, v := range piles {
		if v > maxPiles {
			maxPiles = v
		}
	}

	start, end := 1, maxPiles
	ans := 0

	for start <= end {
		totalDuration := 0 // 29 30
		mid := int((start + end) / 2)
		for _, v := range piles {
			totalDuration += int(math.Ceil(float64(v) / float64(mid)))
		}
		if totalDuration <= h {
			end = mid - 1
			ans = mid
		} else {
			start = mid + 1
		}
	}

	return ans
}
