package Week_08

import "sort"

func merge(intervals [][]int) [][]int {
	//排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	// 合并
	res:= make([][]int, 0)
	cs:= make( []int, 0)
	for _, s:=range intervals {
		if len(res)==0|| cs[1]<s[0] {
			res=append(res,s)
			cs=s
		}else if  cs[1]<s[1]{
			cs[1]=s[1]
		}
	}
	return res
}


