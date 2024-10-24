package main

func searchMatrix(matrix [][]int, target int) bool {
	rows := len(matrix)
	cols := len(matrix[0])
	start, end := 0, rows*cols

	for start < end {
		mid := (start + end) / 2
		r := mid / cols
		c := mid % cols
		if matrix[r][c] == target {
			return true
		}
		if matrix[r][c] > target {
			end = mid
		} else {
			start = mid + 1
		}
	}
	return false
}
