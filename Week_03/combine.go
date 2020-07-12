package main

func combine(n int, k int) [][]int {
	res := make([][]int, 0)
	drill(n, 1, k, &[]int{}, &res)
	return res
}

func drill(n, i, k int, nums *[]int, res *[][]int) {

	if len(*nums) == k {
		tmp := make([]int, 0)
		for _, v := range *nums {
			tmp = append(tmp, v)
		}

		*res = append(*res, tmp)
		return

	}

	for j := i; j <= n; j++ {
		*nums = append(*nums, j)
		drill(n, j+1, k, nums, res)
		//*nums = (*nums)[:len(*nums)-1]
		*nums = (*nums)[:len(*nums)-1]
	}

}
