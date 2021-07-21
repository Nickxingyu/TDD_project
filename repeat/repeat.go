package iteration

const repeat_count int = 5

func Repeat(s string) string {
	var result string
	for i := 0; i < repeat_count; i++ {
		result += s
	}
	return result
}
