package main

import (
	"container/heap"
	"container/list"
)

// https://leetcode.com/problems/kth-largest-sum-in-a-binary-tree/
// BFS + Max Heap

type HeapInt []int64

func (h HeapInt) Len() int           { return len(h) }
func (h HeapInt) Less(i, j int) bool { return h[i] > h[j] }
func (h HeapInt) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *HeapInt) Push(x any) {
	*h = append(*h, x.(int64))
}

func (h *HeapInt) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func kthLargestLevelSum(root *TreeNode, k int) int64 {

	visited := make(map[*TreeNode]bool)
	sums := make([]int64, 0)
	cntLevels := 0

	queue := list.New()
	queue.PushBack(root)
	visited[root] = true

	// Level sum
	for queue.Len() > 0 {
		level := queue.Len()
		sum := int64(0)
		cntLevels++

		for range level {
			element := queue.Front()
			n := element.Value.(*TreeNode)
			sum += int64(n.Val)
			queue.Remove(element)

			if n.Left != nil {
				left := n.Left
				visited[left] = true
				queue.PushBack(left)
			}

			if n.Right != nil {
				right := n.Right
				visited[right] = true
				queue.PushBack(right)
			}
		}

		sums = append(sums, sum)
	}

	if cntLevels < k {
		return -1
	}

	hp := HeapInt(sums)
	heap.Init(&hp)

	var ans int64
	for k > 0 {
		ans = heap.Pop(&hp).(int64)
		k--
	}
	return ans
}
