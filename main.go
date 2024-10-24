package main

import (
	"fmt"
)

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	matrix := [][]int{
		{2, 1, 1},
		{1, 1, 0},
		{0, 1, 1},
	}

	result := orangesRotting(matrix)
	fmt.Println(result)
}
