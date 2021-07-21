package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	assert_correct_message := func(t *testing.T, result string, expect string) {
		t.Helper()
		if result != expect {
			t.Errorf("result: %s  expect: %s", result, expect)
		}
	}
	t.Run("Repeat a char five time", func(t *testing.T) {
		result := Repeat("a", 5)
		expect := "aaaaa"
		assert_correct_message(t, result, expect)
	})
	t.Run("Repeat a char a specific number of times", func(t *testing.T) {
		result := Repeat("b", 10)
		expect := "bbbbbbbbbb"
		assert_correct_message(t, result, expect)
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func ExampleRepeat() {
	result := Repeat("n", 3)
	fmt.Println(result)
	//Output: nnn
}
