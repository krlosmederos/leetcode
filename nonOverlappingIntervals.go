package main

import (
	"sort"
)

// https://leetcode.com/problems/non-overlapping-intervals/description/

func eraseOverlapIntervals(intervals [][]int) int {

	var ans [][]int
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})

	for _, interval := range intervals {
		if len(ans) == 0 {
			ans = append(ans, interval)
		}

		if ans[len(ans)-1][1] <= interval[0] {
			ans = append(ans, interval)
		}
	}

	return len(intervals) - len(ans)
}
