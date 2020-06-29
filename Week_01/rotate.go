package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2}
	k := 3
	rotate(a, k)
	fmt.Println(a)
}

//func rotate(nums []int, k int)  {
//	l := len(nums)
//	if l == 1{
//
//		return
//	}
//	if k > l {
//		k = k % l
//	}
//	swapIndex := 0
//	startIndex := 0
//	for i := 1; i <=  l; i++ {
//		swapIndex = swapIndex + k
//		if swapIndex >= l {
//			swapIndex = swapIndex - l
//		}
//		if startIndex == swapIndex{
//			startIndex ++
//			swapIndex = startIndex
//			continue
//		}
//		nums[startIndex], nums[swapIndex] = nums[swapIndex], nums[startIndex]
//
//	}
//
//}

//反转数组
func rotate(nums []int, k int) {
	l := len(nums)
	reverse(nums, 0, l-1)
	k = k % l
	reverse(nums, 0, k-1)
	reverse(nums, k, l-1)

}

func reverse(nums []int, start int, end int) {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}

}
