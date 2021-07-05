package problems

import (
	model "probsolbygolang/model"
)

// creating a new node with given data
func NewNode(data int) *model.BSTNode {
	var node model.BSTNode
	node.Data = data
	return &node
}

// convert the sorted array to BST
func SortedArrayToBST(arr []int, begin int, end int) *model.BSTNode {

	if begin > end {
		return nil
	}

	// make the middle element as root node
	var mid int = (begin + end) / 2
	root := NewNode(arr[mid])

	// building the left and right subtrees recursively
	root.Left = SortedArrayToBST(arr, begin, mid-1)
	root.Right = SortedArrayToBST(arr, mid+1, end)
	return root
}
