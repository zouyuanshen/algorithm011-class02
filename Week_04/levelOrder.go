package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {

	res := make([][]int, 0)

	dfs(root, 0, &res)
	return res
}

func dfs(root *TreeNode, level int, res *[][]int) {
	if root == nil {
		return
	}
	if len(*res) < level+1{
		*res = append(*res, []int{})
	}
	(*res)[level] = append((*res)[level], root.Val)
	if root.Left !=  nil{
		dfs(root.Left,level+1, res)
	}
	if root.Right !=  nil{
		dfs(root.Right,level+1, res)
	}
}
