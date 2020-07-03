package main

import "container/list"

func main() {

}

var nums []int

func inorderTraversal(root *TreeNode) []int {
	appendInt(root)

	return nums
}

func appendInt(root *TreeNode) {
	if root.Left != nil {
		appendInt(root.Left)
	}
	nums = append(nums, root.Val)
	if root.Right != nil {
		appendInt(root.Right)
	}

}

//stack
func inorderTraversal1(root *TreeNode) []int {
	nums := make([]int, 0)
	cur := root
	//
	stack := list.List{}
	for cur != nil || stack.Len() != 0 {
		for cur != nil {
			stack.PushFront(cur) // 把root的左边所有结点以此推入底部
			cur = cur.Left
		}

		cur = stack.Remove(stack.Front()).(*TreeNode)
		nums = append(nums, cur.Val)
		cur = cur.Right

	}

	return nums
}
