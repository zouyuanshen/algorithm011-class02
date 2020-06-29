package main

import "fmt"

func main() {
	a := []int{0, 0, 1, 2, 0, 3}
	fmt.Println(moveZeros(a))
	//fmt.Println(moveZeros1(a))
}

func moveZeros(nums []int) []int {
	l := len(nums)
	if l <= 1 {
		return nums
	}
	j := 0
	for i := 0; i < l; i++ {
		if nums[i] != 0 {
			nums[j] = nums[i]
			if i != j {
				nums[i] = 0

			}
			j++
		}

	}
	return nums
}

//func moveZeros1(nums []int) []int {
//	rnums := make([]int, 0)
//	for _, v := range nums {
//		if v != 0 {
//			rnums = append(rnums, v)
//		}
//	}
//
//}
