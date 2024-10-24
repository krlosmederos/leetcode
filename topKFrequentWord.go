package main

import (
	"container/heap"
)

// https://leetcode.com/problems/top-k-frequent-words/description/

type node struct {
	word string
	freq int
}

type HeapLex []node

func (h HeapLex) Len() int { return len(h) }
func (h HeapLex) Less(i, j int) bool {
	if h[i].freq == h[j].freq {
		return h[i].word < h[j].word
	}
	return h[i].freq > h[j].freq
}
func (h HeapLex) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *HeapLex) Push(x any) {
	*h = append(*h, x.(node))
}

func (h *HeapLex) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func topKFrequent(words []string, k int) []string {

	seen := make(map[string]int)
	ans := make([]string, 0, k)
	hp := &HeapLex{}

	for _, v := range words {
		seen[v]++
	}

	for k, v := range seen {
		heap.Push(hp, node{word: k, freq: v})
	}

	heap.Init(hp)

	for k > 0 {
		ans = append(ans, (heap.Pop(hp)).(node).word)
		k--
	}

	return ans
}
