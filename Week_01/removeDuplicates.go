package main

import "fmt"

func main() {

	a := []int{1, 1, 2, 2, 3, 3, 4, 4, 4}
	fmt.Println(removeDuplicates1(a))
}

func removeDuplicates(nums []int) int {
	compareIndex := 0
	changeIndex := 0
	// 从index 0 开始 重复则跳跃
	for i, _ := range nums {

		if compareIndex == i {
			continue
		} else {
			if nums[i] == nums[compareIndex] {
				changeIndex = compareIndex + 1

			} else {
				if changeIndex != 0 {
					nums[changeIndex], nums[i] = nums[i], nums[changeIndex]
					compareIndex = changeIndex
					changeIndex++

				} else {
					compareIndex = i
				}
			}
		}

	}
	//fmt.Println(nums)
	return compareIndex + 1
}

func removeDuplicates1(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	j := 0
	if len(nums) == 1 {
		return 1
	}
	for i := j + 1; i < len(nums); i++ {
		if nums[i] != nums[j] {
			j++
			nums[j] = nums[i]
		}
	}
	fmt.Println(nums)
	return j + 1
}
