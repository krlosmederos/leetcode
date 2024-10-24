package main

import "container/heap"

// https://leetcode.com/problems/top-k-frequent-elements/

type HeapQ [][]int

func (h HeapQ) Len() int           { return len(h) }
func (h HeapQ) Less(i, j int) bool { return h[i][1] > h[j][1] }
func (h HeapQ) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *HeapQ) Push(x any) {
	*h = append(*h, x.([]int))
}

func (h *HeapQ) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func topKFrequent1(nums []int, k int) []int {

	seen := make(map[int]int)
	ans := make([]int, 0, k)
	hp := &HeapQ{}

	for _, v := range nums {
		seen[v]++
	}

	for k, v := range seen {
		heap.Push(hp, []int{k, v})
	}

	heap.Init(hp)

	for k > 0 {
		ans = append(ans, (heap.Pop(hp)).([]int)[0])
		k--
	}

	return ans
}
