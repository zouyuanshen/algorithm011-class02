package main

import "fmt"

func main() {
	total := 3
	fmt.Println(climbStairs(4))
	fmt.Println(total)

}

// 尾递归
func climbStairs(n int) int {
	return climbStairsTail(n, 1, 1)
}

func climbStairsTail(n int, n1 int, n2 int) int {
	if n == 0 {
		return n1
	}
	return climbStairsTail(n-1, n2, n1+n2)
}
