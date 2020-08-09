package Week_07

// 分治法
func climbStairs(n int) int {
	i, j := 1, 1
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	for t := 0; t < n-2; t++ {
		i, j = j, i+j
	}
	return j + i
}
