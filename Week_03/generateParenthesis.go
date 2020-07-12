package main

import "fmt"

func main() {
	fmt.Println(generateParenthesis(3))

	s := "dddd"
	fmt.Println(s[:1])
	d := make(map[string]int)
	dictionary := []string{"d"}
	for _, v := range dictionary {
		d[v] = 1
	}
	if d[s[:2]] == 0 {
		fmt.Println(dictionary[1])
	}

}

func generateParenthesis(n int) []string {
	res := make([]string, 0)
	makeParenthesis(0, 0, n, "", &res)

	return res

}

func makeParenthesis(left, right, max int, s string, sarr *[]string) {

	if left == max && left == right {

		*sarr = append(*sarr, s)
		return
	}

	if left < max {
		makeParenthesis(left+1, right, max, s+"(", sarr)
	}
	if right < left {
		makeParenthesis(left, right+1, max, s+")", sarr)
	}
}
