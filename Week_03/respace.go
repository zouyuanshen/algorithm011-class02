package main

func respace(dictionary []string, sentence string) int {
	d := make(map[string]bool)
	for _, v := range dictionary {
		d[v] = true
	}
	n := len(sentence)
	f := make([]int, n+1)
	f[0] = 0
	for i := 1; i < n+1; i++ {

		f[i] = f[i-1] + 1
		for j := 0; j < i; j++ {

			if d[sentence[j:i]] {
				f[i] = min(f[i], f[j])
			}
		}
	}
	return f[len(f)-1]

}

func min(i, j int) int {
	if i > j {
		return j
	} else {
		return i
	}
}
