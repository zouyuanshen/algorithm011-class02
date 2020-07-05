package main

var nums []int

func preorderTraversal(root *TreeNode) []int {

	nums = []int{}
	if root == nil {
		return nums
	}
	appendInt(root)
	return nums
}

func appendInt(root *TreeNode) {

	nums = append(nums, root.Val)
	if root.Left != nil {
		appendInt(root.Left)
	}
	if root.Right != nil {
		appendInt(root.Right)
	}

}
