package sum

import (
	"testing"
)

func TestSum(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	result := Sum(arr)
	want := 15
	if result != want {
		t.Errorf("Sum(1, 2) = %d; want %d", result, want)
	}
}
