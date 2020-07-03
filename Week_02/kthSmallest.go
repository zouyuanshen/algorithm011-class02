package main

import (
	"fmt"
)

func main() {

	matrix := [][]int{{1, 5, 9},
		{10, 11, 13},
		{12, 13, 15}}
	k := 8
	fmt.Println(kthSmallest(matrix, k))
	fmt.Println(1<<3, 2^3)
}

func kthSmallest(matrix [][]int, k int) int {
	//二分查找法 计算 大于和小于中数的
	n := len(matrix)
	left, right := matrix[0][0], matrix[n-1][n-1]
	for left < right {

		mid := (left + right) / 2

		if check(matrix, mid, k, n) {
			right = mid
		} else {

			left = mid + 1
		}

	}

	return left

}

func check(matrix [][]int, mid, k, n int) bool {

	i, j := n-1, 0
	num := 0 // 计算小于mid的个数
	for i >= 0 && j < n {
		if matrix[i][j] <= mid {
			num += i + 1 // 此竖列的全部元素
			j++
		} else {
			i--
		}

	}
	return num >= k
}

func kthSmallest1(matrix [][]int, k int) int {
	m := len(matrix[0])
	n := len(matrix)
	// 先算出第几阶层

	kmod := k % m
	l := (k - kmod) / m
	target := 0
	if n == 1 || l == 0 {
		return matrix[0][k-1]
	}
	for i, j := m-1, m-1; i >= 0 && j >= 0; {
		res := 0
		if matrix[l-1][i] <= matrix[l-2][j] {
			res = matrix[l-2][j]
			j--

		} else {
			res = matrix[l-1][i]
			i--

		}
		target++
		if target == m-kmod {
			return res
		}
	}
	if target < m-kmod {
		return matrix[l-1][m-1-kmod]
	}
	return 0
}
