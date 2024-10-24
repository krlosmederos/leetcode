package main

import (
	"sort"
)

// https://leetcode.com/problems/partition-labels/
func partitionLabels(s string) []int {
	intervals := make([][]int, 0)
	start := make(map[byte]int)
	end := make(map[byte]int)
	merged := make([][]int, 0)
	ans := make([]int, 0)

	for i, v := range s {
		if start[byte(v)] == 0 {
			start[byte(v)] = i + 1
		} else {
			end[byte(v)] = i + 1
		}
	}

	for key := range start {
		if v, found := end[key]; found {
			intervals = append(intervals, []int{start[key], v})
		} else {
			intervals = append(intervals, []int{start[key], start[key]})
		}

	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	for _, interval := range intervals {
		if len(merged) == 0 {
			merged = append(merged, interval)
			continue
		}
		if interval[0] <= merged[len(merged)-1][1] {
			merged[len(merged)-1][1] = max(merged[len(merged)-1][1], interval[1])
			continue
		}

		merged = append(merged, interval)
	}

	for _, v := range merged {
		ans = append(ans, v[1]-v[0]+1)
	}

	return ans
}
