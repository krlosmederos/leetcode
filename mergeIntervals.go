package main

import "sort"

func merge(intervals [][]int) [][]int {
	var merged [][]int

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	j := 0
	for _, interval := range intervals {
		if len(merged) == 0 {
			merged = append(merged, interval)
			j++
			continue
		}

		if interval[0] <= merged[j-1][1] {
			merged[j-1][1] = max(interval[1], merged[j-1][1])
			continue
		}

		merged = append(merged, interval)
		j++
	}
	return merged
}
