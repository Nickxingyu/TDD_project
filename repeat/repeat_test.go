package iteration

import (
	"testing"
)

func TestRepeat(t *testing.T) {
	result := Repeat("a")
	expect := "aaaaa"
	if result != expect {
		t.Errorf("result: %s  expect: %s", result, expect)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}
