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
