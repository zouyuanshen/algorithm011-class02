package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(minimumTotal([][]int{[]int{2}, []int{3, 4}, []int{6, 5, 7}, []int{4, 1, 8, 3}}))
}

// dp
func minimumTotal(triangle [][]int) int {
	n := len(triangle)
	f := make([]int, n)
	f[0] = triangle[0][0]

	for i := 1; i < n; i++ {

		f[i] = f[i-1] + triangle[i][i]

		for j := i - 1; j > 0; j-- {
			f[j] = min(f[j-1], f[j]) + triangle[i][j] // 一定要从后面计算， 因为只有一个一维数组保存状态
		}

		f[0] += triangle[i][0] // 必须最后处理，因为只有一个一维数组保存状态，所以改了f[0]会影响上一层数据
	}
	ans := math.MaxInt32
	for i := 0; i < n; i++ {
		ans = min(ans, f[i])
	}
	return ans
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
