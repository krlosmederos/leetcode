package main

func insert(intervals [][]int, newInterval []int) [][]int {
	var merged, ans [][]int
	inserted := false
	for _, interval := range intervals {
		if newInterval[0] < interval[0] && !inserted {
			merged = append(merged, newInterval)
			inserted = true
		}
		merged = append(merged, interval)
	}
	if !inserted {
		merged = append(merged, newInterval)
	}

	j := 0
	for _, interval := range merged {
		if len(ans) == 0 {
			ans = append(ans, interval)
			j++
			continue
		}
		if interval[0] <= ans[j-1][1] {
			ans[j-1][1] = max(interval[1], ans[j-1][1])
			continue
		}
		ans = append(ans, interval)
		j++
	}

	return ans
}
