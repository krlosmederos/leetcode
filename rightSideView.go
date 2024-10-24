package main

import "container/list"

// https://leetcode.com/problems/binary-tree-right-side-view/description/
// BFS

func rightSideView(root *TreeNode) []int {

	ans := make([]int, 0)
	queue := list.New()
	lvl := 0

	if root == nil {
		return ans
	}

	queue.PushBack(root)

	for queue.Len() > 0 {
		lvl++
		for range queue.Len() {
			n := queue.Front().Value.(*TreeNode)
			if len(ans) < lvl {
				ans = append(ans, n.Val)
			}

			if n.Right != nil {
				queue.PushBack(n.Right)
			}
			if n.Left != nil {
				queue.PushBack(n.Left)
			}
			queue.Remove(queue.Front())
		}

	}

	return ans
}
