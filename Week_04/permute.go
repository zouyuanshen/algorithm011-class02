package main

import (
	"fmt"
)

func main() {
	fmt.Println(permute([]int{1, 2, 3}))
}

func permute(nums []int) [][]int {
	res := make([][]int, 0)
	l := len(nums)
	makeInt(&nums, l, 0, &[]int{}, &res)
	return res

}

func makeInt(cnums *[]int, l, j int, temp *[]int, res *[][]int) {
	// termiation
	if len(*temp) == l {
		t := make([]int, 0)
		t = append(t, *temp...)
		*res = append(*res, t)
		return
	}

	for i := j; i < l; i++ {
		*temp = append(*temp, (*cnums)[i])
		(*cnums)[i], (*cnums)[j] = (*cnums)[j], (*cnums)[i]
		makeInt(cnums, l, j+1, temp, res)
		*temp = (*temp)[:len(*temp)-1]
		(*cnums)[i], (*cnums)[j] = (*cnums)[j], (*cnums)[i]
	}
}
