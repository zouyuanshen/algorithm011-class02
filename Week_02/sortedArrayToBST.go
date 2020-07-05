package main

//
func sortedArrayToBST(nums []int) *TreeNode {
	l := len(nums)
	return setNode(nums, 0, l-1)
}

func setNode(nums []int, left, right int) *TreeNode {

	mid := (left + right) / 2
	root := &TreeNode{
		Val: nums[mid],
	}
	root.Left = setNode(nums, left, mid-1)
	root.Right = setNode(nums, mid+1, right)
	return root
}
