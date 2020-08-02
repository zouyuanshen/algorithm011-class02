package main

func minPathSum(grid [][]int) int {

	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	r, c := len(grid), len(grid[0])
	dp := make([][]int, r)

	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, c)
	}
	dp[0][0] = grid[0][0]
	for i := 1; i < r; i++ {
		dp[i][0] = dp[i - 1][0] + grid[i][0]
	}

	for j := 1; j < c; j++ {
		dp[0][j] = dp[0][j-1] + grid[0][j]
	}
	for i := 1; i< r; i++{
		for j := 1; j < c; j++ {
			dp[i][j] = min(dp[i - 1][j], dp[i][j - 1]) + grid[i][j]
		}
	}
	return dp[r-1][c-1]
}

func min(x, y int)  int{
	if x < y {
		return x
	}
	return y

}