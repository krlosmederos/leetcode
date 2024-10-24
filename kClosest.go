package main

import (
	"container/heap"
	"math"
)

// https://leetcode.com/problems/k-closest-points-to-origin/

type TupleHeap [][]int

func (h TupleHeap) Len() int { return len(h) }

func (h TupleHeap) Less(i, j int) bool {
	ai := h[i][0] * h[i][0]
	bi := h[i][1] * h[i][1]
	ansi := math.Sqrt(float64(ai + bi))
	aj := h[j][0] * h[j][0]
	bj := h[j][1] * h[j][1]
	ansj := math.Sqrt(float64(aj + bj))

	return ansi < ansj
}

func (h TupleHeap) Swap(i, j int) {
	h[i][0], h[j][0] = h[j][0], h[i][0]
	h[i][1], h[j][1] = h[j][1], h[i][1]
}

func (h *TupleHeap) Push(x any) {
	*h = append(*h, x.([]int))
}

func (h *TupleHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func kClosest(points [][]int, k int) [][]int {
	hp := TupleHeap(points)
	heap.Init(&hp)
	ans := make([][]int, 0)

	for k > 0 {
		ans = append(ans, heap.Pop(&hp).([]int))
		k--
	}
	return ans
}
