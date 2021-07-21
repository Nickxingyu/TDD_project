package iteration

func Repeat(s string, number int) string {
	var result string
	repeat_count := number
	for i := 0; i < repeat_count; i++ {
		result += s
	}
	return result
}
