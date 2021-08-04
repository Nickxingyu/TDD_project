package sum

func Sum(arr []int) int {
	var sum int
	for _, v := range arr {
		sum += v
	}
	return sum
}
