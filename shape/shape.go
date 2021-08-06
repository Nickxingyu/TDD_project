package shape

func Perimeter(input []int) int {
	var perimeter int
	for _, v := range input {
		perimeter += v
	}
	return perimeter
}
