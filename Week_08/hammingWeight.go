package Week_08


func hammingWeight(num uint32) int {
	var c int
	for num !=0{
		if num & 1 == 1{
			c++
		}
		num >>= 1//num右移一位
	}
	return c
}
