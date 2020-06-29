package main

import "fmt"

func main() {
	a := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(maxArea(a))
}

func maxArea(nums []int) int {

	max := 0
	i := 0
	j := len(nums) - 1

	for i < j {
		long := 0
		weigth := j - i
		if nums[i] < nums[j] {
			long = nums[i]
			i++
		} else {
			long = nums[j]
			j--
		}
		area := long * weigth
		if area > max {
			max = area
		}
	}
	return max
}
