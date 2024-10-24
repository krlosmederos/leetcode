package main

import (
	"container/list"
)

// https://leetcode.com/problems/rotting-oranges
// BFS matrix

type Tuple struct {
	x int
	y int
}

func orangesRotting(grid [][]int) int {
	queue := list.New()
	fresh, ans := 0, 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
				fresh++
			}
			if grid[i][j] == 2 {
				queue.PushBack(&Tuple{x: i, y: j})
			}
		}
	}

	movs := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

	for queue.Len() > 0 && fresh > 0 {
		ans++
		for range queue.Len() {
			elem := queue.Front()
			x, y := elem.Value.(*Tuple).x, elem.Value.(*Tuple).y
			queue.Remove(elem)

			for _, m := range movs {
				mx, my := m[0]+x, m[1]+y
				if mx >= 0 && mx < len(grid) && my >= 0 && my < len(grid[0]) && grid[mx][my] == 1 {
					nn := &Tuple{x: mx, y: my}
					queue.PushBack(nn)
					fresh--
					grid[mx][my] = 2
				}
			}
		}
	}

	if fresh > 0 {
		return -1
	}

	return ans
}
