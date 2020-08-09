package Week_07

import (
	"strings"
)

func solveNQueens(n int) [][]string {

	// 循环行， 判断pie， na，clo
	res := make([][]string, 0)

	clos := make([]interface{}, 0)
	pies := make([]interface{}, 0)
	nas := make([]interface{}, 0)

	makeQstrings(&res, []interface{}{}, &clos, &pies, &nas, 0, n)

	return res
}

func makeQstrings(res *[][]string, lres []interface{}, clos *[]interface{}, pies *[]interface{}, nas *[]interface{}, level, n int) {
	if level == n {

		*res = append(*res, convertStrings(lres))
	}

	for clo := 0; clo < n; clo++ {
		// 检查Q的合法

		if InList(clo, clos) == true || InList(clo-level, pies) == true || InList(clo+level, nas) == true {
			continue
		}
		*clos = append(*clos, clo)
		*pies = append(*pies, clo-level)
		*nas = append(*nas, clo+level)
		Qstr := makeQstring(n, clo)
		lres = append(lres, Qstr)
		makeQstrings(res, lres, clos, pies, nas, level+1, n)
		removeList(&lres, Qstr)
		removeList(clos, clo)
		removeList(pies, clo-level)
		removeList(nas, clo+level)
	}

}

func InList(i interface{}, li *[]interface{}) bool {
	for _, v := range *li {

		if i == v {
			return true
		}
	}
	return false
}

func makeQstring(n, i int) string {
	res := make([]string, 0)
	for j := 0; j < n; j++ {
		if j == i {
			res = append(res, "Q")
		} else {
			res = append(res, ".")
		}

	}
	s := strings.Join(res, "")

	return s
}

func removeList(li *[]interface{}, value interface{}) {
	index := -1
	for i, v := range *li {
		if v == value {
			index = i
		}
	}
	if index == -1 {
		return
	}
	*li = append((*li)[:index], (*li)[index+1:]...)
}

func convertStrings(li []interface{}) []string {
	strli := make([]string, 0)
	for _, v := range li {

		strli = append(strli, v.(string))
	}
	return strli
}
