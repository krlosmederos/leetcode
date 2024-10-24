package main

import (
	"container/heap"
)

// https://leetcode.com/problems/kth-largest-element-in-an-array

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func findKthLargest(nums []int, k int) int {
	hp := IntHeap(nums)
	heap.Init(&hp)

	var ans any
	for k > 0 {
		ans = heap.Pop(&hp)
		k--
	}
	return ans.(int)
}
