/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 // what are the concepts? 
 // BST - val at left of root is less than root 
 // val at right of root is more than root 
 // to insert we first need to 
 // 1) by convention, it's always easier to insert at the leaf
 // recursive function - how to traverse a tree?
 // how to add a value to a leaf 

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		// we got to a leaf, insert
		return &TreeNode{Val:val}
	}
    //determine if to go left or right based on val
	if val < root.Val {
		root.Left = insertIntoBST(root.Left,val)
	}
	if val > root.Val {
		root.Right = insertIntoBST(root.Right,val)
	}
	return root
}
