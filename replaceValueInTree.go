package main

import "container/list"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func replaceValueInTree(root *TreeNode) *TreeNode {
	sums := make([]int, 0)
	cntLevels := 0

	queue := list.New()
	queue.PushBack(root)

	// Level sum
	for queue.Len() > 0 {
		level := queue.Len()
		sum := 0
		cntLevels++

		for range level {
			element := queue.Front()
			n := element.Value.(*TreeNode)
			sum += n.Val
			queue.Remove(element)

			if n.Left != nil {
				left := n.Left
				queue.PushBack(left)
			}

			if n.Right != nil {
				right := n.Right
				queue.PushBack(right)
			}
		}

		sums = append(sums, sum)
	}

	var newRoot = root
	queue.PushBack(root)
	cntLevels = 0

	// Build the new Binary Tree
	for queue.Len() > 0 {
		level := queue.Len()
		cntLevels++

		for range level {
			element := queue.Front()
			n := element.Value.(*TreeNode)
			queue.Remove(element)

			if n.Left != nil {
				left := n.Left
				queue.PushBack(left)
			}

			if n.Right != nil {
				right := n.Right
				queue.PushBack(right)
			}

			if level == 1 {
				n.Val = 0
			}
			l := 0
			if n.Left != nil {
				l = n.Left.Val
			}
			r := 0
			if n.Right != nil {
				r = n.Right.Val
			}
			if n.Left != nil {
				n.Left.Val = sums[cntLevels] - l - r
			}
			if n.Right != nil {
				n.Right.Val = sums[cntLevels] - l - r
			}

		}
	}

	return newRoot
}
