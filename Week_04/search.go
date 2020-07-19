package main

import "math"

//搜索旋转排序数组

func search(nums []int, target int) int {
	n := len(nums)
	if n == 0 {
		return -1
	}
	l,r,mid := 0 ,n-1, math.MaxInt32
	for l<= r{
		mid = (l+r)>>1
		if nums[mid] == target{
			return mid
		}
		if nums[0] <= nums[mid] {
			if nums[0] <= target && target < nums[mid] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		} else {
			if nums[mid] < target && target <= nums[n - 1] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}

	}
	return -1
}