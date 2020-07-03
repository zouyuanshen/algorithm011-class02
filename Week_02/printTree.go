package main

import "strconv"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func printTree(root *TreeNode) [][]string {
	h := getHeight(root)
	w := (1 << h) - 1
	ans := make([][]string, h)
	fill(root.Left, &ans, 0, 0, w-1, w)
	return ans
}

func fill(root *TreeNode, ans *[][]string, h, s, e, w int) {
	mid := (s + e) / 2
	(*ans)[h] = make([]string, w+1)
	for i := 1; i <= w; i++ {

		(*ans)[h][i] = ""

	}
	(*ans)[h][mid] = strconv.Itoa(root.Val)
	fill(root.Left, ans, h+1, s, mid, w)
	fill(root.Right, ans, h+1, mid+1, e, w)

}

func getHeight(root *TreeNode) int {
	num := 0
	if root == nil {
		return num
	}
	return 1 + Max(getHeight(root.Left), getHeight(root.Right))
}

func Max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}

}

func main() {

}
