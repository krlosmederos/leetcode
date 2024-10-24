package main

import (
	"container/heap"
	"math"
	"slices"
)

// https://leetcode.com/problems/find-k-closest-elements/

type nodee struct {
	val  int
	diff int
}

type HeapNode []nodee

func (h HeapNode) Len() int { return len(h) }
func (h HeapNode) Less(i, j int) bool {
	if h[i].diff == h[j].diff {
		return h[i].val < h[j].val
	}
	return h[i].diff < h[j].diff
}
func (h HeapNode) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *HeapNode) Push(x any) {
	*h = append(*h, x.(nodee))
}

func (h *HeapNode) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func findClosestElements(arr []int, k int, x int) []int {

	ans := make([]int, 0, k)
	hp := &HeapNode{}

	for _, v := range arr {
		n := nodee{
			val:  v,
			diff: int(math.Abs(float64(v - x))),
		}
		heap.Push(hp, n)
	}

	heap.Init(hp)

	for k > 0 {
		ans = append(ans, (heap.Pop(hp)).(nodee).val)
		k--
	}

	slices.Sort(ans)

	return ans
}
