package main

import "sort"

func findMinArrowShots(points [][]int) int {

	merged := make([][]int, 0)

	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})

	for _, point := range points {
		if len(merged) == 0 {
			merged = append(merged, point)
			continue
		}

		if point[0] <= merged[len(merged)-1][1] {
			merged[len(merged)-1][1] = min(point[1], merged[len(merged)-1][1])
			continue
		}

		merged = append(merged, point)
	}

	return len(merged)
}
